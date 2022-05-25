package middlewares

import (
    "Gohub/pkg/jwt"
    "Gohub/pkg/response"

    "github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc{
	return func (c *gin.Context)  {
		if len(c.GetHeader("Authorization")) > 0 {
			//解析token成功, 说明登录成功
			_, err := jwt.NewJWT().ParserToken(c)

			if err == nil {
				response.Unauthorized(c,"请使用游客身份登录")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}