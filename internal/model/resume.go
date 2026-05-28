// Package model defines the data models that are use to represent the
// structure of a resume.
package model

import "net/url"

type Resume struct {
	Contact    ContactInformation `json:"contact"`
	Experience []Employment       `json:"experience"`
	Projects   []Project          `json:"projects"`
	Skills     []Skill            `json:"skills"`
}

type ContactInformation struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     struct {
		ExitCode    *string `json:"exitCode,omitempty"`
		CountryCode *uint8  `json:"countryCode,omitempty"`
		AreaCode    uint16  `json:"areaCode"`
		Prefix      uint16  `json:"prefix"`
		Subscriber  uint16  `json:"subscriber"`
	} `json:"phone"`
	Email    string `json:"email"`
	Website  Link   `json:"website"`
	LinkedIn Link   `json:"linkedin"`
}

type Employment struct {
	Title     string   `json:"title"`
	Employer  string   `json:"employer"`
	Location  *string  `json:"location,omitempty"`
	StartDate string   `json:"startDate"`
	EndDate   *string  `json:"endDate,omitempty"`
	Bullets   []string `json:"bullets"`
	DoRender  bool     `json:"doRender"`
}

type Project struct {
	Title          string   `json:"title"`
	Category       string   `json:"category"`
	Stack          []string `json:"stack"`
	CompletionDate *string  `json:"completionDate,omitempty"`
	URL            *url.URL `json:"url,omitempty"`
	Bullets        []string `json:"bullets"`
	DoRender       bool     `json:"doRender"`
}

type Skill struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

type ProfessionalService struct {
	Title        string  `json:"title"`
	Organization string  `json:"organization"`
	StartDate    string  `json:"startDate"`
	EndDate      *string `json:"endDate,omitempty"`
	Description  string  `json:"description"`
	DoRender     bool    `json:"doRender"`
}
