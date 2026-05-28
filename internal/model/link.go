package model

import "net/url"

type Link struct {
	Title string  `json:"title"`
	Href  url.URL `json:"href"`
}
