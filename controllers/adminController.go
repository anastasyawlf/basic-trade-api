package controllers

import (
	"basic-trade/database"
	"basic-trade/helpers"
	"basic-trade/models/entity"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appJSON = "application/json"
)

func RegisterAdmin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	User := entity.Admin{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	newUUID := uuid.New()
	User.UUID = newUUID.String()

	err := db.Debug().Create(&User).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    User,
	})

}

func LoginAdmin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	User := entity.Admin{}
	var password string

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
