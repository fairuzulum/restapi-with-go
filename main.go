package main

import(
	"go-restapi/controllers"
	"github.com/gin-gonic/gin"
	"go-restapi/models"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}