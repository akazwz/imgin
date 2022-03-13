package model

import (
	"fmt"
	"gorm.io/datatypes"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Image struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;column:id;type:varchar(255);comment: 图片id;"`
	ImageName string    `json:"file_name" gorm:"not null;comment: 图片名;"`
	Album     string    `json:"album" gorm:"not null;default:root;comment:相册;"`
	Unique    string    `json:"unique" gorm:"not null;unique;comment:确保唯一"`
	Size      int64     `json:"size" gorm:"not null;default:0;comment:文件大小;"`
	SHA256    string    `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);comment:文件sha256;"`
	UID       uuid.UUID `json:"uid" gorm:"not null;type:varchar(255);comment:用户uid;"`
	CreatedAt int       `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int       `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (i Image) TableName() string {
	return "image"
}

// BeforeCreate hooks
func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.NewV4()
	i.Unique = fmt.Sprintf("%s-%s-%s", i.UID, i.Album, i.ImageName)
	return
}

type Album struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;column:id;type:varchar(255);comment: 相册id;"`
	Album     string    `json:"album" gorm:"not null;comment:相册;"`
	Unique    string    `json:"unique" gorm:"not null;comment:确保唯一"`
	UID       uuid.UUID `json:"uid" gorm:"not null;type:varchar(255);comment:用户uid;"`
	CreatedAt int       `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int       `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (a Album) TableName() string {
	return "album"
}

// BeforeCreate hooks
func (a *Album) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewV4()
	a.Unique = fmt.Sprintf("%s-%s", a.UID, a.Album)
	return
}

type ImageURI struct {
	ID              uuid.UUID      `json:"id" gorm:"primary_key;column:id;type:varchar(255);comment: id;"`
	SHA256          string         `json:"sha256" gorm:"column:sha256;not null;unique;type:varchar(255);comment:图片sha256;"`
	StorageProvider datatypes.JSON `json:"storage_provider"`
	CreatedAt       int            `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt       int            `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (imageURI ImageURI) TableName() string {
	return "image_uri"
}

// BeforeCreate hooks
func (imageURI *ImageURI) BeforeCreate(tx *gorm.DB) (err error) {
	imageURI.ID = uuid.NewV4()
	return
}
