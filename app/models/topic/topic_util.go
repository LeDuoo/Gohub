package topic

import (
	"Gohub/pkg/app"
	"Gohub/pkg/database"
	"Gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topic Topic) {
	database.DB.Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int, whereMap interface{}) (topics []Topic, paging paginator.Paging) {
	dbQuery := database.DB.Model(Topic{})
	whereArray := whereMap.(map[string]interface{})
	if len(whereArray) > 0 {
		where := ""
		whereAnd := "and"
		if title, ok := whereArray["title"]; ok {
			where = "title like '" + title.(string) + "%'"
		}

		if id, ok := whereArray["id"]; ok {
			if where != "" {
				where = where + whereAnd + " id = " + id.(string)
			} else {
				where = where + " id = " + id.(string)
			}
		}
		dbQuery = dbQuery.Where(where)
	}
	paging = paginator.Paginate(
		c,
		dbQuery,
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
