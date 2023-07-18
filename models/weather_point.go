package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type WeatherPoint struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
	DailyWeathers []DailyWeather `gorm:"foreignkey:WeatherPointID"`
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
	Timezone  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
