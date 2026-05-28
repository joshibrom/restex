package model

import "net/url"

type Link struct {
	Title string
	Href  url.URL
}
