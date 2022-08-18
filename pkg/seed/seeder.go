// Package seed 处理数据库填充相关逻辑
package seed

import (
	"Gohub/pkg/console"
	"Gohub/pkg/database"

	"gorm.io/gorm"
)

type Seeder struct {
	Func SeederFunc
	Name string
}

type SeederFunc func(*gorm.DB)

//按顺序执行的Seeder 数组
//支持一些必须按顺序执行的 seeder, 例如 topic 创建时必须依赖于user,所以TopicSeeder 应该在UserSeeder后执行
var orderedSeederNames []string

//存放所有 Seeder
var seeders []Seeder

//Add 注册到 seeders 数据中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Func: fn,
		Name: name,
	})
}

//SetRunOrder 设置 按顺序执行的Seeder 数组
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有 Seeder
func RunAll() {

	//先运行 ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	//再运行剩下的
	for _, sdr := range seeders {
		//过滤已运行的
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}

}

//RunSeeder 运行单个Seeder
func RunSeeder(name string) {

	for _, sdr := range seeders {
		if name == sdr.Name {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			break
		}
	}
}
