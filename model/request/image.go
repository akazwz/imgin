package request

// UploadImage 上传图片
type UploadImage struct {
	ImageName       string            `json:"image_name" form:"image_name" binding:"required"`
	Album           string            `json:"album" form:"album"`
	Size            int64             `json:"size" form:"size" binding:"required"`
	Sha256          string            `json:"sha256" form:"sha256" binding:"required"`
	StorageProvider []StorageProvider `json:"storage_provider"`
}

type StorageProvider struct {
	Provider string `json:"provider"`
	CID      string `json:"cid"`
}

// NewAlbum  新建相册
type NewAlbum struct {
	AlbumName string `json:"album_name" form:"album_name" binding:"required"`
}
