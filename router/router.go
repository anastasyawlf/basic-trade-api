package router

import (
	"basic-trade/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	adminRouter := router.Group("/auth")
	{
		adminRouter.POST("/register", controllers.RegisterAdmin)
		adminRouter.POST("/login", controllers.LoginAdmin)
	}

	protectedRouter := router.Group("/products")
	{
		// productRouter.GET("/", controllers.GetAllProduct)
		protectedRouter.POST("/", controllers.CreateProduct)
	}

	return router
}

// r.POST("/auth/register", controllers.RegisterAdmin)
// r.POST("/auth/register", controllers.LoginAdmin)
// r.GET("/products", controllers.GetAllProduct)
// r.POST("/products", controllers.CreateProduct)
// r.PUT("/products/:productUUID", controllers.UpdateProduct)
// r.DELETE("/products/:productUUID", controllers.DeleteProduct)
// r.GET("/products/:productUUID", controllers.GetProduct)
// r.GET("/products/variants", controllers.GetAllVariant)
// r.POST("/products/variants", controllers.CreateVariant)
// r.PUT("/products/variants/:variantUUID", controllers.UpdateVariant)
// r.DELETE("/products/:productUUID", controllers.DeleteVariant)
// r.GET("/products/variants/:variantUUID", controllers.GetVariant)
