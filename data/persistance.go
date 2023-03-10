package data

import (
	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/charmbracelet/log"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"

	"gorm.io/gorm"
)

type PersistanceService struct {
	db *gorm.DB
}

func (s *PersistanceService) MigrateSchema() error {
	err := s.db.AutoMigrate(&model.Rating{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.Reading{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.Created{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.AuthorJSON{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.LastModified{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.Type{})
	if err != nil {
		return err
	}
	s.db.AutoMigrate(&model.SourceRecord{})
	if err != nil {
		log.Error(err)
		return err
	}
	s.db.AutoMigrate(&model.Author{})
	if err != nil {
		return err
	}
	return nil
}

func (s *PersistanceService) InsertReadings(readings []model.Reading) error {
	batchSize := 1000
	batches := len(readings) / batchSize

	for i := 0; i <= batches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize

		if end > len(readings) {
			end = len(readings)
		}

		batch := readings[start:end]

		if err := s.db.Create(&batch).Error; err != nil {
			return err
		}
	}

	return nil
}

func NewPersistanceService(db *gorm.DB) *PersistanceService {
	return &PersistanceService{db: db}
}

func (s *PersistanceService) InsertReading(reading model.Reading) error {

	return nil
}

func (s *PersistanceService) InsertRatings(ratings []model.Rating) error {
	batchSize := 1000
	batches := len(ratings) / batchSize

	for i := 0; i <= batches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize

		if end > len(ratings) {
			end = len(ratings)
		}

		batch := ratings[start:end]

		if err := s.db.Create(&batch).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *PersistanceService) InsertAuthors(authors []model.Author) error {
	batchSize := 1000
	batches := len(authors) / batchSize

	for i := 0; i <= batches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize

		if end > len(authors) {
			end = len(authors)
		}

		batch := authors[start:end]

		if err := s.db.Create(&batch).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *PersistanceService) GetTop100() ([]model.Rating, error) {

	return nil, nil
}
