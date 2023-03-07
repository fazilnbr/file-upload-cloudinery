package main

import (
	"file-upload-cloudinary/controllers" //add this

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

// @host localhost:6000
// @BasePath /
// @query.collection.format multi

func main() {
	router := gin.Default()

	//add
	router.POST("/file", controllers.FileUpload())
	router.POST("/remote", controllers.RemoteUpload())

	router.Run("localhost:6000")
}
