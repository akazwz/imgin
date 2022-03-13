package router

import (
	v1 "github.com/akazwz/imgin/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/user/profile", v1.GetUserProfileByToken)
}
