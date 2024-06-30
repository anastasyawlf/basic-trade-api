package controllers

import (
	"basic-trade/database"
	"basic-trade/helpers"
	"basic-trade/models/entity"
	"basic-trade/models/request"

	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	// Get the authenticated user
	userData := ctx.MustGet("userData").(jwt5.MapClaims)
	adminID := uint(userData["id"].(float64))

	var productReq request.ProductRequest
	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that image is provided
	if productReq.ImageURL == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(productReq.ImageURL.Filename)

	// Upload the file
	uploadResult, err := helpers.UploadFile(productReq.ImageURL, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new product
	Product := entity.Product{
		UUID:     uuid.New().String(),
		Name:     productReq.Name,
		ImageURL: uploadResult,
		AdminID:  adminID,
	}

	// Save the product to the database
	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Respond with the created product
	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}
