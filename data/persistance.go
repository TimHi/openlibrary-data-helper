package data

import (
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
	"gorm.io/gorm"
)

type PersistanceService struct {
	db *gorm.DB
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

func (s *PersistanceService) GetTop100() ([]model.Rating, error) {

	return nil, nil
}
