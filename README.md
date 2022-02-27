# googleNews
A go wrapper for the google news rss feed

## Functionality
### TopNews()
Returns the top news based on the configured region and language

### Search(), SearchPeriod(), SearchTimeFrame()
Enables searching for specific keywords and retrieving articles for a given period (last 7 days) or a custom time frame (before, after)

### SearchTopics()
Returns the latest articles for topics such as Business, Health, etc

## Installation
```bash
go get github.com/TomRomeo/googleNews
```

## Getting started

```go
func main() {
	gn := googleNews.New("en", "US")

	articles, err := gn.TopNews()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Current top news for your region:")
	for _, art := range *articles {
		log.Println(art.Title)
	}
}

```
