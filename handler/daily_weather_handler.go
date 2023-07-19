package handler

import (
	"github.com/gin-gonic/gin"
	"gin-sample-app/models"
	"gin-sample-app/database"
	"net/http"
	"log"
	"strconv"
)

type DailyWeatherHandler struct {}

func NewDailyWeatherHandler() *DailyWeatherHandler {
	return &DailyWeatherHandler{}
}

func (dwh *DailyWeatherHandler) GetDailyWeathers(c *gin.Context) {
	var dailyWeathers []models.DailyWeather

	db, err := database.ConnectionOpen()
	if err != nil {
		log.Println(err)
		return
	}

	page, _ := strconv.Atoi(c.Param("page"))
	db.Limit(10).Offset(10 * (page - 1)).Find(&dailyWeathers)

	c.JSON(http.StatusOK, dailyWeathers)
}

func (dwh *DailyWeatherHandler) GetDailyWeatherDetail(c *gin.Context) {
	var dailyWeather models.DailyWeather

	db, err := database.ConnectionOpen()
	if err != nil {
		log.Println(err)
		return
	}
	db.First(&dailyWeather, c.Param("id"))

	c.JSON(http.StatusOK, dailyWeather)
}
