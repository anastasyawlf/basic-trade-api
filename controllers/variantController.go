package controllers

import (
	"basic-trade/database"
	"basic-trade/helpers"
	"basic-trade/models/entity"

	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Variant := entity.Variant{}
	ProductID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Variant)
	} else {
		ctx.ShouldBind(&Variant)
	}

	Variant.ProductID = ProductID
	newUUID := uuid.New()
	Variant.UUID = newUUID.String()

	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}
