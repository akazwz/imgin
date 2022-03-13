package v1

import (
	"github.com/akazwz/imgin/model"
	"github.com/akazwz/imgin/model/request"
	"github.com/akazwz/imgin/model/response"
	"github.com/akazwz/imgin/service"
	"github.com/gin-gonic/gin"
)

// CreateImage 上传图片
func CreateImage(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	var image request.UploadImage

	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.BadRequest(CodeErrorBindJson, "参数错误", c)
		return
	}

	err = service.UploadImageService(image, userUID)
	if err != nil {
		response.BadRequest(CodeErrorCreatFile, "上传失败", c)
		return
	}
	response.Created(CodeSuccessCreateFile, nil, "上传成功", c)
}

// CreateAlbum 新建相册
func CreateAlbum(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	var album request.NewAlbum

	err := c.ShouldBindJSON(&album)
	if err != nil {
		response.BadRequest(CodeErrorBindJson, "参数错误", c)
		return
	}

	err = service.CreateAlbumService(album, userUID)
	if err != nil {
		response.BadRequest(CodeErrorCreateFolder, "新建相册失败", c)
		return
	}
	response.Created(CodeSuccessCreateFolder, nil, "新建成功", c)
}

// GetImagesByAlbum 根据相册获取图片
func GetImagesByAlbum(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	/* 获取文件路径前缀 */
	album := c.Query("album")

	if len(album) < 1 {
		response.BadRequest(CodeErrorParams, "参数错误", c)
		return
	}

	err, fileList := service.GetImagesByAlbumService(userUID, album)
	if err != nil {
		response.BadRequest(CodeErrorGetFileList, "获取图片列表失败", c)
		return
	}

	response.Ok(CodeSuccessGetFileList, fileList, "获取成功", c)
}
