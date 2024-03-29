// Package middlewares Gin 中间件
package middlewares

import (
	"Gohub/app/models/user"
	"Gohub/pkg/config"
	"Gohub/pkg/jwt"
	"Gohub/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

//检测用户是否登录
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claim, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claim.UserID)
		//用户不存在
		if userModel.ID == 0 {
			response.Unauthorized(c, "用户不存在")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
