package controllers

import (
	"file-upload-cloudinary/dtos"
	"file-upload-cloudinary/models"
	"file-upload-cloudinary/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary upload image from users local dierectory
// @ID upload-local-file
// @Produce json
// @Param file formData file true "image"
// @Success 200 {object} dtos.MediaDto{}
// @Failure 500 {object} dtos.MediaDto{}
// @Router /file [post]
func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formFile, _, err := c.Request.FormFile("file")
		fmt.Printf("\n%v\n", formFile)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}

// @Summary upload image from users remote dierectory
// @ID upload-remote-file
// @Produce json
// @Param url body models.Url{} true "url of the image"
// @Success 200 {object} dtos.MediaDto{}
// @Failure 500 {object} dtos.MediaDto{}
// @Router /remote [post]
func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url models.Url

		//validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				dtos.MediaDto{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
