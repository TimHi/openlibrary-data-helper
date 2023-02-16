package model

import "time"

// Data struct for an entry from the Rating dump: https://openlibrary.org/developers/dumps
type Rating struct {
	WorkKey    string
	EditionKey string
	Rating     int
	Date       time.Time
}