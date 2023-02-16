package data

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timhi/openlibrary-data-helper/m/v2/database"
	"github.com/timhi/openlibrary-data-helper/m/v2/model"
)

// Service define a service
type PersistanceService struct {
	r *database.Queries
}

// NewService cria um novo serviço. Lembre-se: receba interfaces, retorne structs ;)
func NewPersistanceService(r *database.Queries) *PersistanceService {
	return &PersistanceService{
		r: r,
	}
}

func (s *PersistanceService) InsertReading(ctx context.Context, reading model.Reading) error {

	err := s.r.InsertReading(ctx, database.InsertReadingParams{
		Workkey: reading.WorkKey,
		Editionkey: sql.NullString{
			String: reading.EditionKey,
			Valid:  true,
		},
		Shelf: sql.NullString{
			String: reading.Shelf,
			Valid:  true,
		},
		Datestamp: sql.NullString{
			String: reading.Date.String(),
			Valid:  true,
		},
	})

	if err != nil {
		return fmt.Errorf("error creating reading: %w", err)
	}

	return nil
}

func (s *PersistanceService) InsertRating(ctx context.Context, rating model.Rating) error {

	err := s.r.InsertRating(ctx, database.InsertRatingParams{
		Workkey: rating.WorkKey,
		Editionkey: sql.NullString{
			String: rating.EditionKey,
			Valid:  true,
		},
		Ratingvalue: sql.NullFloat64{
			Float64: rating.Rating,
			Valid:   true,
		},
		Datestamp: sql.NullString{
			String: rating.Date.String(),
			Valid:  true,
		},
	})

	if err != nil {
		return fmt.Errorf("error creating rating: %w", err)
	}

	return nil
}

func (s *PersistanceService) GetTop100(ctx context.Context) ([]model.Rating, error) {
	top100Data := []model.Rating{}
	top100RawData, err := s.r.ListTop100(ctx)

	if err != nil {
		return nil, fmt.Errorf("error creating rating: %w", err)
	}

	for _, row := range top100RawData {
		convertedDate, _ := time.Parse("YYYY-MM-DD", row.Datestamp.String)
		top100Data = append(top100Data, model.Rating{WorkKey: row.Workkey, EditionKey: row.Editionkey.String, Rating: row.Ratingvalue.Float64, Date: convertedDate})
	}

	return top100Data, nil
}
