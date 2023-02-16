package model

import "time"

// Data struct for an entry from the reading dump: https://openlibrary.org/developers/dumps
// Work Key,   		Edition Key (optional), Shelf, Date
// /works/OL63060W	/books/OL5816906M	Already Read	2017-12-26
type Reading struct {
	WorkKey    string
	EditionKey string
	Shelf      string
	Date       time.Time
}
