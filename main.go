package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"log"

	"github.com/timhi/openlibrary-data-helper/m/v2/data"
	"github.com/timhi/openlibrary-data-helper/m/v2/database"
	"github.com/timhi/openlibrary-data-helper/m/v2/parser"
)

//go:embed data/schema.sql
var ddl string

// go run main.go -reading /Users/hiller/dev/openlibrary-data-helper/dumps/reading.txt
func main() {
	var readingLocation = flag.String("reading", "", "location for the reading data dump")
	var ratingLocation = flag.String("rating", "", "location for the rating data dump")
	flag.Parse()

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Panic(err)
	}

	queries := database.New(db)
	persistanceService := data.NewPersistanceService(queries)

	if *readingLocation != "" {
		log.Println("Start parsing reading data dump...")
		err := parser.ReadingData(*readingLocation, persistanceService, ctx)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("No reading file path specified")
	}

	if *ratingLocation != "" {
		log.Println("Start parsing ratingLocation data dump...")
		err := parser.RatingData(*ratingLocation, persistanceService, ctx)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("No file path specified")
	}
	log.Println("Everything done!")
}
