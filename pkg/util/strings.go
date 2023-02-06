package util

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/gunrosen/search-engine/internal/model"
)

func RefineEditorJsText(str string) string {
	str = strings.TrimPrefix(str, "\"")
	str = strings.TrimSuffix(str, "\"")
	str = strings.Replace(str, "\\", "", -1)
	str = strings.Replace(str, "\u00a0", "", -1) // remove &nbsp
	str = TrimHTML(str)
	return str
}

func MarshalEditorJs(str string) model.EditorJsModel {
	str = RefineEditorJsText(str)
	var jsonData model.EditorJsModel
	err := json.Unmarshal([]byte(str), &jsonData)
	if err != nil {
		println(err)
	}
	return jsonData
}

func TrimHTML(str string) string {
	re := regexp.MustCompile(`<[^>]*(>|$)`)
	trimStr := re.ReplaceAllString(str, "")
	re = regexp.MustCompile(`&nbsp;|&zwnj;|&raquo;|&laquo;|&gt;|\n`)
	return re.ReplaceAllString(trimStr, " ")
}
