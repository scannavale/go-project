package controllers

import (
	"github.com/gin-gonic/gin"
	"go-project/database"
	"go-project/models"
	"net/http"
	"strconv"
)

type PostBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetAllPosts(ctx *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	ctx.JSON(http.StatusOK, posts)
}

func CreatePost(ctx *gin.Context) {
	var body PostBody

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{"Unprocessable body"})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}
	database.DB.Create(&post)
	ctx.JSON(http.StatusCreated, post)
}

func UpdatePost(ctx *gin.Context) {
	var body PostBody
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{"Unprocessable body"})
		return
	}

	var post models.Post

	if result := database.DB.Where("id = ?", uint(id)).First(&post); result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{"Post not found"})
		return
	}

	post.Title, post.Content = body.Title, body.Content
	database.DB.Save(&post)
	ctx.JSON(http.StatusOK, post)
}

func DeletePost(ctx *gin.Context)  {
	var post models.Post
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if result := database.DB.Where("id = ?", uint(id)).First(&post); result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{"Post not found"})
		return
	}

	database.DB.Delete(&post)
	ctx.JSON(http.StatusOK, post)
}