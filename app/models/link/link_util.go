package link

import (
	"Gohub/pkg/app"
	"Gohub/pkg/cache"
	"Gohub/pkg/database"
	"Gohub/pkg/helpers"
	"Gohub/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	//设置缓存
	cacheKey := "links:all"
	//设置过期时间
	expireTime := 120 * time.Minute
	//查询缓存
	cache.GetObject(cacheKey, &links)

	//缓存内无数据
	if helpers.Empty(links) {
		links = All()
		if !helpers.Empty(links) {
			//将查询数据存储到缓存
			cache.Set(cacheKey, links, expireTime)
		}
	}
	return links
}
