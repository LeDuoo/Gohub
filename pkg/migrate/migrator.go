// Package migrate 处理数据库迁移
package migrate

import (
	"Gohub/pkg/console"
	"Gohub/pkg/database"
	"Gohub/pkg/file"
	"io/ioutil"

	"gorm.io/gorm"
)

//Migrate 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration  对应数据的migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;uniquer;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例, 用以执行迁移操作
func NewMigrator() *Migrator {

	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	//migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}

	//不存在才创建
	if !migrator.Migrator.HasTable(&migration) { //GORM 约定使用结构体名的复数形式作为表名
		migrator.Migrator.CreateTable(&migration)
	}
}

// Up执行所有未迁移过的文件
func (migrator *Migrator) Up() {

	//读取所有迁移文件,确保按照时间排序
	migrateFiles := migrator.readAllMigrationFiles()

	// 获取当前批次的值
	batch := migrator.getBatch()

	//获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations) //查询migrations表所有数据

	// 可以通过此值来判断数据库是否已是最新
	runed := false
	// 对迁移文件进行遍历, 如果没有执行过, 就执行up 回调
	for _, mfile := range migrateFiles {
		// 对比文件名称, 看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up data. ")
	}
}

// Rollback 回滚会后一次迁移
func (migrator *Migrator) Rollback() {

	//获取最后一批次迁移数据的批次号
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	//根据最后批次号获取最后一批次所有数据
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	//回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] tabe is empty, nothing to rollback")
	}
}

// Reset 回滚所有迁移
func (migrator *Migrator) Reset() {

	migrations := []Migration{}

	//按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)

	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset")
	}
}

//Refresh 回滚所有迁移, 并运行所有迁移
func (migrator *Migrator) Refresh() {

	//回滚所有迁移
	migrator.Reset()

	//再次执行所有迁移文件
	migrator.Up()
}

//从文件目录读取文件, 保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {

	//读取database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)
	var migrateFiles []MigrationFile
	for _, f := range files {

		// 去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())

		// 通过迁移文件的名词获取 MigrationFile 对象
		mfile := getMigrationFile(fileName)

		// 判断迁移文件是否可用, 再放进migrateFiles 数组中
		if len(mfile.FileName) > 0 {

			migrateFiles = append(migrateFiles, mfile)
		}
	}

	// 返回排序好的 MigrationFile 数组
	return migrateFiles

}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {

	// 默认为1
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	//如果有值的话 加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

//执行迁移 ,执行迁移的up方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {

	//执行 up区块 的SQL
	if mfile.Up != nil {
		//友好提示
		console.Warning("migrating " + mfile.FileName)
		//执行 up 方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}

	//入库
	err := migrator.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error
	console.ExitIf(err)
}

// 回退迁移, 按照倒序执行迁移的 down方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {

	//标记是否真的有执行了迁移回退的操作
	runed := false

	for _, _migration := range migrations {

		//友好提示
		console.Warning("rollback " + _migration.Migration)

		// 执行迁移文件的down方法
		mfile := getMigrationFile(_migration.Migration) //通过迁移文件的名称来获取到 MigrationFile对象, 再由对象进行回滚或迁移操作

		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true

		//回退成功了就删除这条记录
		migrator.DB.Delete(&_migration)
		// 打印运行状态
		console.Success("finsh " + mfile.FileName)
	}

	return runed
}

// Fresh Dorp 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {

	//获取数据库名称,用以提示
	dbname := database.CurrentDatabase()

	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Success("clearup database " + dbname)

	//重新创建 migrates表
	migrator.createMigrationsTable()
	console.Success("[migrations] table crteated.")

	//重新调用 up 命令
	migrator.Up()
}
