package request

import "mime/multipart"

type ProductRequest struct {
	Name     string                `form:"product_name" binding:"required"`
	ImageURL *multipart.FileHeader `form:"image_url" binding:"required"`
}
