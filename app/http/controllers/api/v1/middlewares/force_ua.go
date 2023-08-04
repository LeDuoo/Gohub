// Package middlewares Gin 中间件
package middlewares

import (
	"Gohub/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

//ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.UserAgent() == "" {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}
