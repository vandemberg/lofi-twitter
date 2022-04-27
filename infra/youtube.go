package infra

import (
	"flag"
	"log"
	"net/http"
	config "twitter-lofi/config"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func Youtube(searchText string) []*youtube.SearchResult {
	flag.Parse()
	developerKey := config.GoDotEnvVariable("GOOGLE_API_KEY")

	var (
		query          = flag.String("q", searchText, "Search text")
		maxResults     = flag.Int64("max-results", 50, "Max YouTube results")
		publishedAfter = flag.String("publishedAfter", "2022-01-01T00:00:00Z", "format RFC 3339")
	)

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	listParams := []string{"id", "snippet"}
	call := service.Search.List(listParams).
		Q(*query).
		MaxResults(*maxResults).
		PublishedAfter(*publishedAfter).
		SafeSearch("strict").
		Type("video")

	response, err := call.Do()

	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	return response.Items
}
