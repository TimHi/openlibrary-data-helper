package data

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

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
		Ratingvalue: sql.NullInt64{
			Int64: int64(rating.Rating),
			Valid: true,
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
