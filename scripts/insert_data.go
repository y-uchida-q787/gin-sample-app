package main

import (
	"encoding/json"
	"gin-sample-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"github.com/shopspring/decimal"
	"time"
)

type Daily struct {
	Time []string `json:"time"`
	Temperature2mMax []decimal.Decimal `json:"temperature_2m_max"`
	Temperature2mMin []decimal.Decimal `json:"temperature_2m_min"`
	Weathercode []int `json:"weathercode"`
}

type WeatherApiData struct {
	Latitude decimal.Decimal
	Longitude decimal.Decimal
	Timezone string
	Daily Daily
}

func createNewsArticles() {
	// NewsAPI の APIにリクエストして 「大谷翔平」に関すニュースを取得する
	url := "https://newsapi.org/v2/everything?q=%E5%A4%A7%E8%B0%B7%E7%BF%94%E5%B9%B3&from=2023-07-01&sortBy=publishedAt&apiKey=f1871965ddc848cbaac48e37c52a10fa"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// JSONを構造体にエンコード
	var response map[string][]map[string]string
	json.Unmarshal(body, &response)

	dsn := "root:password@tcp(127.0.0.1:3306)/gin_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	for _, v := range response["articles"] {
		article := models.NewsArticle{
			Title:        v["title"],
			Description:  v["description"],
			Content:      v["content"],
			ArticleUrl:   v["url"],
			ImageUrl:     v["urlToImage"],
			ResourceName: v["author"],
		}

		db.Create(&article)
	}
}

func createWeatherInfo() {
	// open-meteo のAPIにアクセスして天気予報のデータを取得する
	url := "https://api.open-meteo.com/v1/forecast?latitude=35.6895&longitude=139.6917&daily=weathercode,temperature_2m_max,temperature_2m_min&timezone=Asia%2FTokyo"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// JSONを構造体にエンコード
	var response WeatherApiData
	json.Unmarshal(body, &response)

	dsn := "root:password@tcp(127.0.0.1:3306)/gin_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	latitude := response.Latitude
	longitude := response.Longitude

  point := models.WeatherPoint{
		Latitude:    latitude,
		Longitude:   longitude,
		Timezone:    response.Timezone,
	}
	db.Create(&point)

	for i := 0; i < len(response.Daily.Time) ; i++ {
		weatherDate, _ := time.Parse("2006-01-02", response.Daily.Time[i])

		daily := models.DailyWeather{
			WeatherPointID:  point.ID,
			WeatherCode:   response.Daily.Weathercode[i],
			MaxTemperature: response.Daily.Temperature2mMax[i],
			MinTemperature: response.Daily.Temperature2mMin[i],
			WeatherDate:    weatherDate,
		}
		db.Create(&daily)
	}
}

func main() {
	createNewsArticles()
	createWeatherInfo()
}
