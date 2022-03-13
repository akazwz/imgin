package cloud

import (
	"os"

	v1 "github.com/akazwz/imgin/api/v1"
	"github.com/akazwz/imgin/model/response"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
)

// GetAliCloudSTS 获取阿里云 STS
func GetAliCloudSTS(c *gin.Context) {
	/* 获取 阿里云 ak as region */
	accessKeyId := os.Getenv("ALI_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALI_ACCESS_KEY_SECRET")
	regionId := os.Getenv("ALI_REGION_ID")

	/* 阿里云 sts */
	stsClient, err := sts.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = os.Getenv("ALI_ASSUME_ROLE_REQUEST_SCHEME")
	request.RoleArn = os.Getenv("ALI_ASSUME_ROLE_REQUEST_ROLE_ARN")
	request.RoleSessionName = "oss-uploader"

	res, err := stsClient.AssumeRole(request)

	if err != nil {
		response.BadRequest(v1.CodeErrorCreateAliSTS, "获取阿里云STS失败", c)
		return
	}

	ak := res.Credentials.AccessKeyId
	as := res.Credentials.AccessKeySecret
	stk := res.Credentials.SecurityToken

	type AliSTS struct {
		AK  string `json:"ak"`
		AS  string `json:"as"`
		STK string `json:"stk"`
	}

	response.Ok(v1.CodeSuccessGetAliSTS, AliSTS{
		AK:  ak,
		AS:  as,
		STK: stk,
	}, "获取阿里云STS成功", c)
}
