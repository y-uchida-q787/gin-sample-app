package server

import "github.com/gin-gonic/gin"

func NewsServerStart() {
	router := gin.New()

	// 引数が無い場合は、デフォルトの設定で PORT: 8080 で起動する
	router.Run(":8080")
}

func WeatherServerStart() {
	router := gin.New()

	// 引数が無い場合は、デフォルトの設定で PORT: 8080 で起動する
	router.Run(":8091")
}
