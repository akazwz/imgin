package router

import (
	"github.com/akazwz/imgin/api/v1/cloud"
	"github.com/akazwz/imgin/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(routerGroup *gin.RouterGroup) {
	/* 注册用户 */
	routerGroup.POST("/user", user.CreateUserByUsernamePwd)
	/* 登录获取token */
	routerGroup.POST("/user/token", user.CreateTokenByUsernamePwd)
	routerGroup.GET("/cloud/alists", cloud.GetAliCloudSTS)
}
