package handler

import (
	"github.com/gin-gonic/gin"
	"gin-sample-app/models"
	"gin-sample-app/database"
	"net/http"
	"log"
	"strconv"
)

type NewsArticleHandler struct {}

func NewNewsArticleHandler() *NewsArticleHandler {
	return &NewsArticleHandler{}
}

func (nah *NewsArticleHandler) GetNewsArticles(c *gin.Context) {
	var newsArticles []models.NewsArticle

	db, err := database.ConnectionOpen()
	if err != nil {
		log.Println(err)
		return
	}

	page, _ := strconv.Atoi(c.Param("page"))
	db.Limit(10).Offset(10 * (page - 1)).Find(&newsArticles)

	c.JSON(http.StatusOK, newsArticles)
}

func(nah *NewsArticleHandler) GetNewsArticleDetail(c *gin.Context) {
	var newsArticle models.NewsArticle

	db, err := database.ConnectionOpen()
	if err != nil {
		log.Println(err)
		return
	}
	db.First(&newsArticle, c.Param("id"))

	c.JSON(http.StatusOK, newsArticle)
}
