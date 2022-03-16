package router

import (
	"github.com/akazwz/imgin/api/v1/cloud"
	"github.com/gin-gonic/gin"
)

func InitCloudRouter(group *gin.RouterGroup) {
	cloudRouter := group.Group("cloud")
	{
		cloudRouter.GET("aliyun/sts", cloud.GetAliCloudSTS)
		cloudRouter.GET("qiniu/upload-token", cloud.GetUploadToken)
	}
}
