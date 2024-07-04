package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/TheusLab/ASN-Project/backend/utils"
	"github.com/elastic/go-elasticsearch/v8"
)

type SearchResult struct {
	Hits struct {
		Hits []struct {
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

var es *elasticsearch.Client

func init() {
	var err error
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func Search(query string) ([]json.RawMessage, error) {
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("documents"),
		es.Search.WithBody(strings.NewReader(`{"query": {"match": {"content": "`+query+`"}}}`)),
	)
	if err != nil {
		utils.Log.Error().Err(err).Msg("Elasticsearch search error")
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response: %s", res.String())
	}

	var result SearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		utils.Log.Error().Err(err).Msg("Elasticsearch decode error")
		return nil, err
	}

	sources := make([]json.RawMessage, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		sources[i] = hit.Source
	}

	return sources, nil
}
