package model

import "gorm.io/gorm"

// https://openlibrary.org/works/OL45804W.json
type Work struct {
	gorm.Model
	Title   string `json:"title"`
	Key     string `json:"key"`
	Authors []struct {
		Author struct {
			gorm.Model
			Key string `json:"key"`
		} `json:"author"`
		Type struct {
			gorm.Model
			Key string `json:"key"`
		} `json:"type"`
	} `json:"authors"`
	Type struct {
		gorm.Model
		Key string `json:"key"`
	} `json:"type"`
	Description    string   `json:"description"`
	Covers         []int    `json:"covers"`
	SubjectPlaces  []string `json:"subject_places"`
	Subjects       []string `json:"subjects"`
	SubjectPeople  []string `json:"subject_people"`
	SubjectTimes   []string `json:"subject_times"`
	Location       string   `json:"location"`
	LatestRevision int      `json:"latest_revision"`
	Revision       int      `json:"revision"`
	Created        struct {
		gorm.Model
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		gorm.Model
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}
