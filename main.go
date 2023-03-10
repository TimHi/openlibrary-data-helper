package main

import (
	_ "embed"
	"flag"

	"github.com/charmbracelet/log"
	"github.com/timhi/openlibrary-data-helper/m/v2/data"

	"github.com/timhi/openlibrary-data-helper/m/v2/parser"
	"github.com/timhi/openlibrary-data-helper/m/v2/transform"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go run main.go -reading /Users/hiller/dev/openlibrary-data-helper/dumps/reading.txt
// go run main.go -author /Users/hiller/dev/openlibrary-data-helper/dumps/authors.txt
// go run main.go -transform top100
func main() {
	var readingLocation = flag.String("reading", "", "path to the reading data dump .txt")
	var authorLocation = flag.String("author", "", "path to the author data dump .txt")
	var ratingLocation = flag.String("rating", "", "path to the rating data dump .txt")
	var transformOperation = flag.String("transform", "", "operation to apply on the data, available options: [top100]")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{Logger: logger.Default, SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal(err)
	}

	persistanceService := data.NewPersistanceService(db)
	err = persistanceService.MigrateSchema()
	if err != nil {
		log.Fatal(err)
	}

	if *readingLocation != "" {
		log.Info("Start parsing reading data dump...")
		err := parser.ReadingData(*readingLocation, persistanceService)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Info("No reading file path specified")
	}

	if *ratingLocation != "" {
		log.Info("Start parsing ratingLocation data dump...")
		err := parser.RatingData(*ratingLocation, persistanceService)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("No file path specified")
	}

	if *authorLocation != "" {
		log.Info("Start parsing the author data dump...")
		err := parser.AuthorData(*authorLocation, persistanceService)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("No file path specified")
	}

	transform.Start(*transformOperation, *persistanceService)
	log.Info("Everything done!")
}
