package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type DailyWeather struct {
	// gorm.Model
	ID             uint         `gorm:"primary_key"`
	WeatherPointID uint
	WeatherCode    int
	MaxTemperature decimal.Decimal
	MinTemperature decimal.Decimal
	WeatherDate    time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
