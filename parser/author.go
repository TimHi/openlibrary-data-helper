package parser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/timhi/openlibrary-data-helper/m/v2/data"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
	"github.com/timhi/openlibrary-data-helper/m/v2/util"
)

func AuthorData(filePath string, persistanceService *data.PersistanceService) error {
	log.Info("Reading lines...")
	authors, err := readAuthorsFromFile(filePath)

	if err != nil {
		return err
	}
	return bulkInsertAuthors(authors, persistanceService)
}

func readAuthorsFromFile(filePath string) ([]model.Author, error) {
	authors := []model.Author{}
	lines, err := util.ReadLines(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(lines))

	for _, line := range lines {
		if line != "" {
			author, err := parseLineToAuthor(line)
			if err != nil {
				//Guess we can ignore one off errors
				log.Error(err)
			} else {
				authors = append(authors, author)
			}
		}
	}

	return authors, nil
}

func parseLineToAuthor(line string) (model.Author, error) {
	///type/author	/authors/OL10000116A	1	2021-12-26T21:31:53.952079	{"type": {"key": "/type/author"}, "name": "Hubert Briechle", "key": "/authors/OL10000116A", "source_records": ["bwb:9788490159781"], "latest_revision": 1, "revision": 1, "created": {"type": "/type/datetime", "value": "2021-12-26T21:31:53.952079"}, "last_modified": {"type": "/type/datetime", "value": "2021-12-26T21:31:53.952079"}}
	fields := strings.Split(line, "\t")
	author := model.Author{}
	authorJSON := &model.AuthorJSON{}

	if err := json.Unmarshal([]byte(fields[4]), authorJSON); err != nil {
		log.Error(err)
		fmt.Println(line)
		return author, err
	}

	createdTime := time.Time{}
	if authorJSON.CreatedAt.Value != "" {
		tCreatedTime, err := time.Parse(time.RFC3339, authorJSON.CreatedAt.Value)
		if err != nil {
			log.Error(err)
			fmt.Println(line)
			return author, err
		}
		createdTime = tCreatedTime
	}

	lastModifiedTime := time.Time{}
	if authorJSON.LastModifiedAt.Value != "" {
		tLastModifiedTime, err := time.Parse(time.RFC3339, authorJSON.LastModifiedAt.Value)
		if err != nil {
			log.Error(err)
			return author, err
		}
		lastModifiedTime = tLastModifiedTime
	}

	author = model.Author{
		AuthorType: model.AuthorType{
			Key: authorJSON.AuthorType.Key,
		},
		Name:           authorJSON.Name,
		Key:            authorJSON.Key,
		LatestRevision: authorJSON.LatestRevision,
		Revision:       authorJSON.Revision,
		Created: model.Created{
			Typ:   authorJSON.CreatedAt.Typ,
			Value: createdTime.Format("2006-01-02 15:04:05"),
		},
		LastMod: model.LastModified{
			Typ:   authorJSON.LastModifiedAt.Typ,
			Value: lastModifiedTime.Format("2006-01-02 15:04:05"),
		},
	}

	for _, sourceRecord := range authorJSON.SourceRecords {
		author.SourceRecords = append(author.SourceRecords, model.SourceRecord{
			SourceRecord: sourceRecord.SourceRecord,
		})
	}

	return author, nil
}

func bulkInsertAuthors(authors []model.Author, persistanceService *data.PersistanceService) error {
	log.Info(fmt.Sprintf("Insert %d authors... \n", len(authors)))
	return persistanceService.InsertAuthors(authors)
}
