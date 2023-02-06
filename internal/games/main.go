package games

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gunrosen/search-engine/internal/model"
)

func CreateMapping(esClient *elasticsearch.Client) error {
	indexName := model.IndexNameGame
	res, err := esClient.Indices.Create(
		indexName,
		esClient.Indices.Create.WithBody(strings.NewReader(model.Game)),
	)
	if err != nil {
		return err
	}
	var mapResp map[string]interface{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&mapResp)
	if err != nil {
		return err
	}
	if mapResp["error"] != nil {
		errStr := fmt.Sprintf("status=%d: index:%s %s \n", int(mapResp["status"].(float64)), indexName, mapResp["error"].(map[string]interface{})["reason"])
		return errors.New(errStr)
	}
	return nil
}
