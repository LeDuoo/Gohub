package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// // 初始化 Gin 实例
	// r := gin.Default()

	// // 注册一个路由
	// r.GET("/", func(c *gin.Context) {

	// 	// 以 JSON 格式响应
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Hello": "World!",
	// 		"XB":    "PQ",
	// 	})
	// })

	// //运行服务
	// r.Run()

	//new 一个 Gin Engine 的实例
	r := gin.New()

	//注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
			"XB":    "PQ",
		})
	})

	//处理 404 请求
	r.NoRoute(func(c *gin.Context) {
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

	//运行服务 指定端口
	r.Run(":8000")

}