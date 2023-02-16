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
	"github.com/timhi/openlibrary-data-helper/m/v2/transform"
)

//go:embed data/schema.sql
var ddl string

// go run main.go -reading /Users/hiller/dev/openlibrary-data-helper/dumps/reading.txt
// go run main.go -transform top100
func main() {
	var readingLocation = flag.String("reading", "", "path to the reading data dump .txt")
	var ratingLocation = flag.String("rating", "", "path to the rating data dump .txt")
	var transformOperation = flag.String("transform", "", "operation to apply on the data, available options: [top100]")
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

	transform.Start(*transformOperation, *persistanceService, ctx)
	log.Println("Everything done!")
}
