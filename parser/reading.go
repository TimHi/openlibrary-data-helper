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
)

func ReadingData(filePath string, persistanceService *data.PersistanceService, ctx context.Context) error {
	log.Println("Reading lines...")
	readings, err := readReadingsFromFile(filePath)

	if err != nil {
		return err
	}
	return bulkInsertReadings(readings, persistanceService, ctx)
}

func bulkInsertReadings(readings []model.Reading, persistanceService *data.PersistanceService, ctx context.Context) error {
	log.Printf("Insert %d readings... \n", len(readings))
	for _, reading := range readings {
		err := persistanceService.InsertReading(ctx, reading)
		if err != nil {
			return err
		}
	}
	return nil
}

func readReadingsFromFile(filePath string) ([]model.Reading, error) {
	readings := []model.Reading{}
	file, err := util.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reading, err := parseLineToReading(scanner.Text())
		if err != nil {
			//Guess we can ignore one off errors
			log.Println(err)
		} else {
			readings = append(readings, reading)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return readings, nil
}

func parseLineToReading(line string) (model.Reading, error) {
	fields := strings.Split(line, "\t")
	if len(fields) < 3 {
		return model.Reading{}, fmt.Errorf("invalid input: %s", line)
	}

	var editionKey string
	if len(fields) > 3 {
		editionKey = fields[1]
	}

	date, err := time.Parse("2006-01-02", fields[len(fields)-1])
	if err != nil {
		return model.Reading{}, fmt.Errorf("invalid date format: %s", fields[len(fields)-1])
	}

	reading := model.Reading{
		WorkKey:    fields[0],
		EditionKey: editionKey,
		Shelf:      fields[2],
		Date:       date,
	}

	return reading, nil
}
