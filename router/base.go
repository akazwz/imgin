package router

import (
	"github.com/akazwz/imgin/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(routerGroup *gin.RouterGroup) {
	/* 注册用户 */
	routerGroup.POST("/user", user.CreateUserByUsernamePwd)
	/* 登录获取token */
	routerGroup.POST("/user/token", user.CreateTokenByUsernamePwd)
}
