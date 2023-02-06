package guilds

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gunrosen/search-engine/internal"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gunrosen/search-engine/internal/model"
)

func CreateMapping(esClient *elasticsearch.Client) error {
	indexName := model.IndexNameGuild
	res, err := esClient.Indices.Create(
		indexName,
		esClient.Indices.Create.WithBody(strings.NewReader(model.GuildMappingType)),
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
	guilds, err := getListInternal()
	if err != nil {
		return err
	}
	var r []model.IndexData
	for _, guild := range guilds {
		staffs := make([]map[string]string, 0)
		for _, staff := range guild.Staffs {
			if staff.Name != "" {
				staffs = append(staffs, map[string]string{
					"name": staff.Name,
				})
			}
		}
		guildData := map[string]interface{}{
			"slug":           guild.Slug,
			"name":           guild.Name,
			"introduction":   guild.Introduction,
			"regions":        guild.Regions,
			"member_regions": guild.MemberRegions,
			"languages":      guild.Languages,
			"difference":     guild.Difference,
			"staffs":         staffs,
		}

		r = append(r, model.IndexData{
			DocumentID: guild.ID,
			Data:       guildData,
		})
	}
	return internal.DoIndex(esClient, model.IndexNameGuild, r)
}

func getListInternal() ([]*model.Guild, error) {
	url := fmt.Sprintf("%s/v1/guilds", os.Getenv("CORE_GAME_URL"))
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
		return nil, fmt.Errorf("guilds: Get error when get internal list %s", url)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	guildRes := new(struct {
		Data struct {
			Items []*model.Guild `json:"items"`
		} `json:"data"`
	})
	if res.StatusCode == http.StatusOK {
		json.Unmarshal(body, guildRes)
	}

	return guildRes.Data.Items, nil
}
