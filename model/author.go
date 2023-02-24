package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID             int `gorm:"primaryKey"`
	AuthorType     Type
	Name           string
	Key            string
	SourceRecord   []SourceRecords
	LatestRevision int
	Revision       int
	CreatedStruct  Created
	LastMod        LastModified
}

type Type struct {
	AuthorID int
	Key      string
}

type LastModified struct {
	AuthorID int
	Typ      string
	Value    string
}

type Created struct {
	ID       int
	AuthorID int
	Typ      string
	Value    string
}

type SourceRecords struct {
	gorm.Model
	AuthorID     int
	SourceRecord string
}

type AuthorJSON struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Name           string   `json:"name"`
	Key            string   `json:"key"`
	SourceRecords  []string `json:"source_records"`
	LatestRevision int      `json:"latest_revision"`
	Revision       int      `json:"revision"`
	Created        struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}
