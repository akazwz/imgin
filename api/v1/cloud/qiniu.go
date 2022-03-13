package cloud

import (
	"os"

	v1 "github.com/akazwz/imgin/api/v1"
	"github.com/akazwz/imgin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// GetUploadFileToken 获取文件上传token(qiniu)
func GetUploadFileToken(c *gin.Context) {
	accessKey := os.Getenv("QINIU_ACCESS_KEY")
	secretKey := os.Getenv("QINIU_SECRET_KEY")

	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "akazwz"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	type TokenUpload struct {
		Token string `json:"token"`
	}

	response.Ok(v1.CodeSuccessCreateUploadToken, TokenUpload{
		Token: upToken,
	}, "获取成功", c)
}
