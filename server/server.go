package server

import (
	"github.com/gin-gonic/gin"
	"gin-sample-app/handler"
	"net/http"
)

func NewBasicAuthServer() http.Handler {
	router := gin.New()

  authorized := router.Group("/user", gin.BasicAuth(gin.Accounts{
    "admin":   "p@ssw0rd",
    "user1":   "password1",
    "user2":   "password2",
  }))

  authorized.GET("/autheticate", func(c *gin.Context) {
    user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, gin.H{ "message": user + " is authenticated!" })
  })

	return router
}

func NewsServer() http.Handler {
	router := gin.New()
	newsArticleHandler := handler.NewNewsArticleHandler()

	v1 := router.Group("/v1")
	{
		v1.GET("/news/:page", newsArticleHandler.GetNewsArticles)
		v1.GET("/news/article/:id", newsArticleHandler.GetNewsArticleDetail)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/news/:page", newsArticleHandler.GetNewsArticles)
		v2.GET("/news/article/:id", newsArticleHandler.GetNewsArticleDetail)
	}

	return router
}

func WeatherServer() http.Handler  {
	router := gin.New()
	dailyWeatherHandler := handler.NewDailyWeatherHandler()

	v1 := router.Group("/v1")
	{
		v1.GET("/weather/:page", dailyWeatherHandler.GetDailyWeathers)
		v1.GET("/weather/daily/:id", dailyWeatherHandler.GetDailyWeatherDetail)
	}

	// 引数が無い場合は、デフォルトの設定で PORT: 8080 で起動する
	return router
}
