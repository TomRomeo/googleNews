package models

import "encoding/xml"

type RssRes struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel represents a collection of google news items
type Channel struct {
	Generator     string    `xml:"generator"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Language      string    `xml:"language"`
	WebMaster     string    `xml:"webMaster"`
	Copyright     string    `xml:"copyright"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Description   string    `xml:"description"`
	Items         []Article `xml:"item"`
}

// Article represents a google news item
type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Guid        string   `xml:"guid"`
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
	Source      string   `xml:"source"`
}
