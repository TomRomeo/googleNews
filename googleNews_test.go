package googleNews

import (
	"github.com/TomRomeo/googleNews/pkg/models"
	"strings"
	"testing"
	"time"
)

func TestGoogleNews_Search(t *testing.T) {

	cl := New("en", "US")

	articles, err := cl.Search("test")
	if err != nil {
		t.Errorf("Search() returned an error: %s", err)
	}

	if len(*articles) == 0 {
		t.Error("Search() returned no articles")
	}

}

func TestGoogleNews_SearchPeriod(t *testing.T) {

	cl := New("en", "US")

	articles, err := cl.SearchPeriod("test", "1d")
	if err != nil {
		t.Errorf("SearchPeriod() returned an error: %s", err)
	}

	if len(*articles) == 0 {
		t.Error("SearchPeriod() returned no articles")
	}

	for _, article := range *articles {

		articleDate := strings.Split(article.PubDate, ", ")[1]
		artDate, err := time.Parse("02 Jan 2006 15:04:05 GMT", articleDate)
		if err != nil {
			t.Error(err)
		}
		if artDate.Before(time.Now().Add(-24 * time.Hour)) {
			t.Error("SearchPeriod returned an article that is younger than the specified period")
		}
	}

}

func TestGoogleNews_SearchTimeframe(t *testing.T) {

	cl := New("en", "US")

	articles, err := cl.SearchTimeframe("test", "2020-06-02", "2020-06-04")
	if err != nil {
		t.Errorf("SearchTimeframe() returned an error: %s", err)
	}
	article0 := (*articles)[0]

	if len(*articles) == 0 {
		t.Error("SearchTimeframe() returned no articles")
	}

	if article0.Title != "Covid-19 testing at home: The race to make it easy as a pregnancy test - Vox.com" {
		t.Error("Article has wrong title. Expected: 'Covid-19 testing at home: The race to make it easy as a pregnancy test - Vox.com', got:", article0.Title)
	}
	for _, article := range *articles {

		articleDate := strings.Split(article.PubDate, ", ")[1]
		artDate, err := time.Parse("02 Jan 2006 15:04:05 GMT", articleDate)
		if err != nil {
			t.Error(err)
		}
		if artDate.After(time.Date(2020, 06, 05, 0, 0, 0, 0, time.UTC)) {
			t.Error("SearchTimeFrame returned an article that is older than the specified before date")
		}
		if artDate.Before(time.Date(2020, 06, 02, 0, 0, 0, 0, time.UTC)) {
			t.Error("SearchTimeFrame returned an article that is younger than the specified after date")
		}
	}

}

func TestGoogleNews_SearchTopic(t *testing.T) {

	cl := New("en", "US")

	articles, err := cl.SearchTopic(models.BusinessTopic)
	if err != nil {
		t.Errorf("SearchTopic() returned an error: %s", err)
	}

	if len(*articles) == 0 {
		t.Error("SearchTopic() returned no articles")
	}

}

func TestGoogleNews_TopNews(t *testing.T) {

	cl := New("en", "US")

	articles, err := cl.TopNews()
	if err != nil {
		t.Errorf("TopNews() returned an error: %s", err)
	}

	if len(*articles) == 0 {
		t.Error("TopNews() returned no articles")
	}

}

func TestGoogleNews_languageAndRegionUrl(t *testing.T) {

	cl := New("en", "US")

	url := cl.languageAndRegionUrl()

	if url != "hl=en&gl=US&ceid=US:en" {
		t.Error("languageAndRegionUrl() returned wrong string")
	}

}

func TestNew(t *testing.T) {
	cl := New("en", "US")

	if cl == nil {
		t.Errorf("New() did not return a GoogleNew instance")
	}

	if cl.Lang != "en" {
		t.Errorf("New() did not set the language parameter correctly")
	}

	if cl.Region != "US" {
		t.Errorf("New() did not set the region parameter correctly")
	}

}
