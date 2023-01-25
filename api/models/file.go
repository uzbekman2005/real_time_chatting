package models

import "mime/multipart"

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadPhotoRes struct {
	URL  string `json:"photo_url"`
	Type string `json:"media_type"`
}

