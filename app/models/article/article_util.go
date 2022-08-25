package article

import (
	"Gohub/pkg/app"
	"Gohub/pkg/database"
	"Gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (article Article) {
	database.DB.Where("id", idstr).First(&article)
	return
}

func GetBy(field, value string) (article Article) {
	database.DB.Where("? = ?", field, value).First(&article)
	return
}

func All() (articles []Article) {
	database.DB.Find(&articles)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Article{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int, whereMap interface{}) (articles []Article, paging paginator.Paging) {

	//原始db查询
	dbQuery := database.DB.Model(Article{})
	whereData := whereMap.(map[string]string)
	if len(whereData) > 0 {
		// dbQuery = database.DB.Where(whereData).First(&articles)//可多条件查询   查询规则为 =
		dbQuery = database.DB.Where("title like ?", whereData["title"]+"%").First(&articles) //模糊查询 需多条件时用and 例 where(user_id = ? and item_name like ?", userId, title+"%)
	}

	paging = paginator.Paginate(
		c,
		dbQuery,
		&articles,
		app.V1URL(database.TableName(&Article{})),
		perPage,
	)
	return
}
