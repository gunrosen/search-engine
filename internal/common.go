package internal

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gunrosen/search-engine/internal/model"
)

func DoIndex(esClient *elasticsearch.Client, indexName string, data []model.IndexData) error {
	for _, d := range data {
		jsonBytes, err := json.Marshal(d.Data)
		if err != nil {
			return err
		}

		req := esapi.IndexRequest{
			Index:      indexName,
			DocumentID: d.DocumentID,
			Body:       bytes.NewReader(jsonBytes),
			Refresh:    "true",
		}
		// Perform the request with the client.
		res, err := req.Do(context.Background(), esClient)
		if err != nil {
			return err
		}
		res.Body.Close()
	}
	return nil
}
