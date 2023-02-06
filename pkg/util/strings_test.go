package util_test

import (
	"encoding/json"
	"testing"

	"github.com/gunrosen/search-engine/internal/model"
	"github.com/gunrosen/search-engine/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestRefineEditorJsText(t *testing.T) {
	in := "{\"time\":1652945586994,\"blocks\":[{\"id\":\"NBQOaZQbtq\",\"type\":\"paragraph\",\"data\":{\"text\":\"ASPO World is going to be an exemplary of the play to earn model (P2E model) in this era of digital assets. Players in ASPO World can engage in the daily activities and PvP battles to claim items, whose values are determined by its rarity. They can then trade them on the Marketplace to earn money.\"}}],\"version\":\"2.23.2\"}"
	out := util.RefineEditorJsText(in)
	var jsonData model.EditorJsModel
	err := json.Unmarshal([]byte(out), &jsonData)
	assert.Equal(t, nil, err, "This text should be able to convert to JSON")
}

func TestMarshalEditorJs(t *testing.T) {
	in := "{\"time\":1652945586994,\"blocks\":[{\"id\":\"NBQOaZQbtq\",\"type\":\"paragraph\",\"data\":{\"text\":\"<b>ASPO World</b>&nbsp; is going to be an exemplary of the play to earn model (P2E model) in this era of digital assets. Players in ASPO World can engage in the daily activities and PvP battles to claim items, whose values are determined by its rarity. They can then trade them on the Marketplace to earn money.\"}}],\"version\":\"2.23.2\"}"
	out := util.MarshalEditorJs(in)
	assert.Equal(t, "ASPO World  is going to be an exemplary of the play to earn model (P2E model) in this era of digital assets. Players in ASPO World can engage in the daily activities and PvP battles to claim items, whose values are determined by its rarity. They can then trade them on the Marketplace to earn money.", out.Blocks[0].Data.Text, "This text should be parsed")
}

func TestTrimHTML(t *testing.T) {
	in := "<b>1. PvP Brawls:</b>&nbsp;Do you have what it takes to beat down those boisterous, boozy robots?\n Go head-to-head with other players to climb the leaderboards throughout each PvP season. At the end of each season, players will be rewarded based on their division and rating, with the best players receiving the biggest prizes.&nbsp;"
	out := util.TrimHTML(in)
	assert.Equal(t, "1. PvP Brawls: Do you have what it takes to beat down those boisterous, boozy robots?  Go head-to-head with other players to climb the leaderboards throughout each PvP season. At the end of each season, players will be rewarded based on their division and rating, with the best players receiving the biggest prizes. ", out)
}
