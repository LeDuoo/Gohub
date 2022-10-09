// Package policies 用户授权
package policies

import (
	"Gohub/app/models/topic"
	"Gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanmodifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
