// Package model defines the data models that are use to represent the
// structure of a resume.
package model

import "net/url"

type Resume struct {
	Contact        ContactInformation
	WorkExperience []Employment
	Projects       []Project
	Skills         []Skill
}

type ContactInformation struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Website   Link
	LinkedIn  Link
}

type Employment struct {
	Title     string
	Employer  string
	Location  *string
	StartDate string
	EndDate   *string
	Bullets   []string
	DoRender  bool
}

type Project struct {
	Title          string
	Category       string
	Stack          []string
	CompletionDate *string
	URL            *url.URL
	Bullets        []string
	DoRender       bool
}

type Skill struct {
	Category string
	Items    []string
}

type ProfessionalService struct {
	Title        string
	Organization string
	StartDate    string
	EndDate      *string
	Description  string
	DoRender     bool
}
