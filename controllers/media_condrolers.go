package controllers

import (
	"file-upload-cloudinary/dtos"
	"file-upload-cloudinary/models"
	"file-upload-cloudinary/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
