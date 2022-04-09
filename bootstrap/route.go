package bootstrap

import (
	"Gohub/app/http/controllers/api/v1/middlewares"
	"Gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	//注册全局中间件
	registerGlobalMiddleWare(router)

	//注册 APi 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	//处理404请求
	router.NoRoute(func(c *gin.Context) {
		// c.Request 是 gin 封装的请求对象，所有用户的请求信息，都可以从这个对象中获取
		//获取标头信息 Accept 信息
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			//如果是html格式
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			//默认返回JSON格式
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    "404",
				"error_massage": "路由未定义,请确认Url和请求方式",
			})
		}
	})
}
