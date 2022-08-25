//Package article 模型
package article

import (
	"Gohub/app/models"
	"Gohub/pkg/database"
)

type Article struct {
	models.BaseModel

	// Put fields in here
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	models.CommonTimestampsField
}

func (article *Article) Create() {
	database.DB.Create(&article)
}

func (article *Article) Save() (rowsAffected int64) {
	result := database.DB.Save(&article)
	return result.RowsAffected
}

func (article *Article) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&article)
	return result.RowsAffected
}
