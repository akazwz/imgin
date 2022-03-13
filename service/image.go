package service

import (
	"github.com/akazwz/imgin/global"
	"github.com/akazwz/imgin/model"
	"github.com/akazwz/imgin/model/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UploadImageService 保存图片信息到数据库
func UploadImageService(uploadImage request.UploadImage, uid uuid.UUID) (err error) {
	imageUri := model.ImageURI{
		SHA256:          uploadImage.Sha256,
		StorageProvider: nil,
	}

	image := model.Image{
		ImageName: uploadImage.ImageName,
		Album:     uploadImage.Album,
		Size:      uploadImage.Size,
		SHA256:    uploadImage.Sha256,
		UID:       uid,
	}

	/* transaction */
	err = global.GORMDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&imageUri).Error
		if err != nil {
			return err
		}
		err = tx.Create(&image).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// CreateAlbumService 保存相册信息到数据库
func CreateAlbumService(newAlbum request.NewAlbum, uid uuid.UUID) (err error) {
	album := model.Album{
		Album: newAlbum.AlbumName,
		UID:   uid,
	}

	err = global.GORMDB.Create(&album).Error
	return
}

// GetImagesByAlbumService 根据相册和uid获取文件列表
func GetImagesByAlbumService(uid uuid.UUID, album string) (err error, images []model.Image) {
	err = global.GORMDB.Where("uid = ? AND album = ?", uid, album).Find(&images).Error
	return
}
