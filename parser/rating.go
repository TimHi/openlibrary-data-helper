package parser

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/timhi/openlibrary-data-helper/m/v2/data"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
	"github.com/timhi/openlibrary-data-helper/m/v2/util"
	"github.com/timhi/swiss-army-knife/src/stringutil"
)

func RatingData(filePath string, persistanceService *data.PersistanceService, ctx context.Context) error {
	log.Println("Reading lines...")
	ratings, err := readRatingsFromFile(filePath)

	if err != nil {
		return err
	}
	return bulkInsertRatings(ratings, persistanceService, ctx)
}

func readRatingsFromFile(filePath string) ([]model.Rating, error) {
	readings := []model.Rating{}
	file, err := util.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rating, err := parseLineToRating(scanner.Text())
		if err != nil {
			//Guess we can ignore one off errors
			log.Println(err)
		} else {
			readings = append(readings, rating)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return readings, nil
}

func parseLineToRating(line string) (model.Rating, error) {
	fields := strings.Split(line, "\t")
	if len(fields) < 3 {
		return model.Rating{}, fmt.Errorf("invalid input: %s", line)
	}

	var editionKey string
	if len(fields) > 3 {
		editionKey = fields[1]
	}

	date, err := time.Parse("2006-01-02", fields[len(fields)-1])
	if err != nil {
		return model.Rating{}, fmt.Errorf("invalid date format: %s", fields[len(fields)-1])
	}

	ratingValue := stringutil.ParseFloat64(fields[2])

	rating := model.Rating{
		WorkKey:    fields[0],
		EditionKey: editionKey,
		Rating:     ratingValue,
		Date:       date,
	}

	return rating, nil
}

func bulkInsertRatings(ratings []model.Rating, persistanceService *data.PersistanceService, ctx context.Context) error {
	log.Printf("Insert %d ratings... \n", len(ratings))
	for _, rating := range ratings {
		err := persistanceService.InsertRating(ctx, rating)
		if err != nil {
			return err
		}
	}
	return nil
}
