package model

import (
	"time"

	"gorm.io/gorm"
)

// Data struct for an entry from the Rating dump: https://openlibrary.org/developers/dumps
type Rating struct {
	gorm.Model
	WorkKey    string
	EditionKey string
	Rating     float64
	Date       time.Time
}
