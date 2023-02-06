package games

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gunrosen/search-engine/internal"

	"github.com/gunrosen/search-engine/pkg/util"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gunrosen/search-engine/internal/model"
)

func CreateMapping(esClient *elasticsearch.Client) error {
	indexName := model.IndexNameGame
	res, err := esClient.Indices.Create(
		indexName,
		esClient.Indices.Create.WithBody(strings.NewReader(model.GameMappingType)),
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

func IngestData(esClient *elasticsearch.Client) error {
	games, err := getListInternal()
	if err != nil {
		return err
	}
	var r []model.IndexData
	for _, game := range games {
		gameData := map[string]interface{}{
			"slug":               game.Slug,
			"name":               game.Name,
			"excerpt":            game.Excerpt,
			"description":        util.MarshalEditorJs(game.Description),
			"roadmap_text":       util.MarshalEditorJs(game.RoadmapText),
			"white_paper":        util.MarshalEditorJs(game.WhitePaper),
			"introduction":       util.MarshalEditorJs(game.Introduction),
			"features_text":      util.MarshalEditorJs(game.FeaturesText),
			"play_mode":          util.MarshalEditorJs(game.PlayMode),
			"play_to_earn_model": util.MarshalEditorJs(game.PlayToEarnModel),
			"guideline":          util.MarshalEditorJs(game.Guideline),
		}

		r = append(r, model.IndexData{
			DocumentID: game.ID,
			Data:       gameData,
		})
	}
	return internal.DoIndex(esClient, model.IndexNameGame, r)
}

func getListInternal() ([]*model.Game, error) {
	url := fmt.Sprintf("%s/v1/games", os.Getenv("CORE_GAME_URL"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("games: Get error when get internal list %s", url)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	gameRes := new(struct {
		Data struct {
			Items []*model.Game `json:"items"`
		} `json:"data"`
	})
	if res.StatusCode == http.StatusOK {
		json.Unmarshal(body, gameRes)
	}

	return gameRes.Data.Items, nil
}
