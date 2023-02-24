package parser

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/timhi/openlibrary-data-helper/m/v2/data"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
	"github.com/timhi/openlibrary-data-helper/m/v2/util"
)

func AuthorData(filePath string, persistanceService *data.PersistanceService) error {
	log.Info("Reading lines...")
	_, err := readAuthorsFromFile(filePath)

	if err != nil {
		return err
	}
	//return bulkInsertAuthors(authors, persistanceService)
	return nil
}

func readAuthorsFromFile(filePath string) ([]model.Author, error) {
	authors := []model.Author{}
	file, err := util.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		author, err := parseLineToAuthor(scanner.Text())
		if err != nil {
			//Guess we can ignore one off errors
			log.Error(err)
		} else {
			authors = append(authors, author)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func parseLineToAuthor(line string) (model.Author, error) {
	///type/author	/authors/OL10000116A	1	2021-12-26T21:31:53.952079	{"type": {"key": "/type/author"}, "name": "Hubert Briechle", "key": "/authors/OL10000116A", "source_records": ["bwb:9788490159781"], "latest_revision": 1, "revision": 1, "created": {"type": "/type/datetime", "value": "2021-12-26T21:31:53.952079"}, "last_modified": {"type": "/type/datetime", "value": "2021-12-26T21:31:53.952079"}}
	fields := strings.Split(line, "\t")
	if len(fields) < 5 {
		return model.Author{}, fmt.Errorf("invalid input: %s", line)
	}
	/*
		var editionKey string
		if len(fields) > 3 {
			editionKey = fields[1]
		}

		date, err := time.Parse("2006-01-02", fields[len(fields)-1])
		if err != nil {
			return model.Author{}, fmt.Errorf("invalid date format: %s", fields[len(fields)-1])
		}

		ratingValue := stringutil.ParseFloat64(fields[2])

		author := model.Author{
			ID:             0,
			AuthorType:     model.Type{},
			Name:           "",
			Key:            editionKey,
			SourceRecord:   []model.SourceRecords{},
			LatestRevision: 0,
			Revision:       0,
			CreatedStruct: model.Created{
				ID:       0,
				AuthorID: 0,
				Typ:      "",
				Value:    "",
			},
			LastMod: model.LastModified{
				AuthorID: 0,
				Typ:      "",
				Value:    "",
			},
		}
	*/
	return model.Author{}, nil
}
