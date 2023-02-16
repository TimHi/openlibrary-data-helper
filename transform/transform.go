package transform

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/timhi/openlibrary-data-helper/m/v2/data"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
)

func Start(operation string, persistanceService data.PersistanceService, ctx context.Context) error {
	var err error
	switch operation {
	case "top100":
		err = Top100AsJson(persistanceService, ctx)
	default:
		log.Println("No transformation specified")
	}
	return err
}

func Top100AsJson(persistanceService data.PersistanceService, ctx context.Context) error {
	top100Ratings, err := persistanceService.GetTop100(ctx)
	top100Works := []model.Work{}
	if err != nil {
		return err
	}

	log.Println("Requesting additional information from the API...")
	for _, rating := range top100Ratings {
		apiItem, err := GetWorkFromApi(rating)
		if err != nil {
			continue
		}
		top100Works = append(top100Works, apiItem)
	}

	return outputToJson(top100Works)
}

func outputToJson(top100Works []model.Work) error {
	log.Println("Writing the responses to json...")
	file, err := os.Create("top100.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := json.Marshal(top100Works)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.Write(data); err != nil {
		log.Fatal(err)
	}
	return err
}
