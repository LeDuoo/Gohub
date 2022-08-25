package migrations

import (
	"Gohub/app/models"
	"Gohub/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Article struct {
		models.BaseModel

		Title   string `gorm:"type:varchar(255);not null;index"`
		Content string `gorm:"type:text;default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Article{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Article{})
	}

	migrate.Add("2022_08_25_142656_add_articles_table", up, down)
}
