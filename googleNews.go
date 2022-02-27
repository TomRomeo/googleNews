package googleNews

import (
	"encoding/xml"
	"fmt"
	"github.com/TomRomeo/googleNews/pkg/models"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

// A GoogleNews struct represents a connection to google news
//
// If you want to initialize a GoogleNews client,
// please use googleNews.New()
type GoogleNews struct {
	Lang    string
	Region  string
	baseUrl string
}

// New returns a new GoogleNews Client
func New(lang, region string) *GoogleNews {
	return &GoogleNews{
		Lang:    lang,
		Region:  region,
		baseUrl: "https://news.google.com/rss/",
	}
}

// A helperfunction to add language and region to urls
func (c *GoogleNews) languageAndRegionUrl() string {
	return fmt.Sprintf("hl=%s&gl=%s&ceid=%s:%s", c.Lang, c.Region, c.Region, c.Lang)
}

// SearchTopic queries the /topics endpoint for the given topic
func (c *GoogleNews) SearchTopic(topic models.Topic) (*[]models.Article, error) {
	url := fmt.Sprintf("%stopics/%s?%s", c.baseUrl, string(topic), c.languageAndRegionUrl())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	var xmlRes models.RssRes

	if err := xml.Unmarshal(data, &xmlRes); err != nil {
		return nil, err
	}
	articles = xmlRes.Channel.Items

	return &articles, nil
}

// SearchPeriod enables searching for a query in a given time period
// example: period: 7d - search query in articles of the last 7 days
func (c *GoogleNews) SearchPeriod(searchTerm string, period string) (*[]models.Article, error) {
	searchTerm = url2.QueryEscape(searchTerm)

	url := fmt.Sprintf("%ssearch?q=%s+when:%s&%s", c.baseUrl, searchTerm, period, c.languageAndRegionUrl())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	var xmlRes models.RssRes

	if err := xml.Unmarshal(data, &xmlRes); err != nil {
		return nil, err
	}
	articles = xmlRes.Channel.Items

	return &articles, nil
}

// SearchTimeframe is similar to SearchPeriod, making it possible to search between two dates
// dates should have the format: 2020-06-02
func (c *GoogleNews) SearchTimeframe(searchTerm string, after, before string) (*[]models.Article, error) {
	searchTerm = url2.QueryEscape(searchTerm)

	url := fmt.Sprintf("%ssearch?q=%s+after:%s+before:%s&%s", c.baseUrl, searchTerm, after, before, c.languageAndRegionUrl())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	var xmlRes models.RssRes

	if err := xml.Unmarshal(data, &xmlRes); err != nil {
		return nil, err
	}
	articles = xmlRes.Channel.Items

	return &articles, nil
}

// Search let's you retrieve articles from a custom search query
func (c *GoogleNews) Search(searchTerm string) (*[]models.Article, error) {
	searchTerm = url2.QueryEscape(searchTerm)

	url := fmt.Sprintf("%ssearch?q=%s&%s", c.baseUrl, searchTerm, c.languageAndRegionUrl())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	var xmlRes models.RssRes

	if err := xml.Unmarshal(data, &xmlRes); err != nil {
		return nil, err
	}
	articles = xmlRes.Channel.Items

	return &articles, nil
}

// TopNews returns the top articles based on the configured region and language
func (c *GoogleNews) TopNews() (*[]models.Article, error) {

	url := fmt.Sprintf("%snews?%s", c.baseUrl, c.languageAndRegionUrl())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	var xmlRes models.RssRes

	if err := xml.Unmarshal(data, &xmlRes); err != nil {
		return nil, err
	}
	articles = xmlRes.Channel.Items

	return &articles, nil
}
