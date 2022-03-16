package cloud

import (
	"os"

	v1 "github.com/akazwz/imgin/api/v1"
	"github.com/akazwz/imgin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// GetUploadToken 获取文件上传token(qiniu)
func GetUploadToken(c *gin.Context) {
	accessKey := os.Getenv("QINIU_ACCESS_KEY")
	secretKey := os.Getenv("QINIU_SECRET_KEY")

	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "akazwz"
	/* 上传策略： 只能上传图片，最大为 10MB,检测内容 mime type */
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		MimeLimit:  "image/*",
		FsizeLimit: 1024 * 1024 * 10,
		DetectMime: 1,
	}
	upToken := putPolicy.UploadToken(mac)

	type TokenUpload struct {
		Token string `json:"token"`
	}

	response.Ok(v1.CodeSuccessCreateUploadToken, TokenUpload{
		Token: upToken,
	}, "获取成功", c)
}
