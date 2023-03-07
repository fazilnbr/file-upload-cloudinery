package main

import (
	"file-upload-cloudinary/controllers" //add this
	_ "file-upload-cloudinary/docs"
	_ "file-upload-cloudinary/dtos"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Go + Gin File Upload API
// @version 1.0
// @description This is a sample server file upload server. You can visit the GitHub repository at https://github.com/fazilnbr/file-upload-cloudinery

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email fazilkp2000@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	router := gin.Default()
	// docs route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// upload route
	router.POST("/file", controllers.FileUpload())
	router.POST("/remote", controllers.RemoteUpload())

	router.Run("localhost:8000")
}
