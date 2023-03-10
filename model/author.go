package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	AuthorType     AuthorType
	Name           string
	Key            string
	SourceRecords  []SourceRecord
	LatestRevision int
	Revision       int
	Created        Created
	LastMod        LastModified
}

type AuthorType struct {
	gorm.Model
	Key string
}

type AuthorJSON struct {
	gorm.Model
	AuthorID       int
	AuthorType     AuthorType
	Name           string
	Key            string
	SourceRecords  []SourceRecordJSON
	LatestRevision int
	Revision       int
	CreatedAt      Created      `gorm:"embedded;embeddedPrefix:created_"`
	LastModifiedAt LastModified `gorm:"embedded;embeddedPrefix:last_modified_"`
}

type SourceRecord struct {
	gorm.Model
	AuthorID     int // added foreign key field
	SourceRecord string
}

type SourceRecordJSON struct {
	SourceRecord string
}

type Type struct {
	Key string
}

type LastModified struct {
	Typ   string
	Value string
}

type Created struct {
	Typ   string
	Value string
}
