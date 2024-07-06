package utils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic/v7"
)

type SearchResult struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func Search(query string) []SearchResult {
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	searchResult, err := client.Search().
		Index("asn").
		Query(elastic.NewQueryStringQuery(query)).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	var results []SearchResult
	for _, hit := range searchResult.Hits.Hits {
		var result SearchResult
		err := json.Unmarshal(hit.Source, &result)
		if err != nil {
			log.Printf("Error unmarshalling document: %s", err)
		}
		results = append(results, result)
	}
	return results
}
