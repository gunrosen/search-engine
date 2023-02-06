package model

type IndexData struct {
	Data       map[string]interface{}
	DocumentID string
}

type EditorJsModel struct {
	Blocks []EditorJsModelBlocks `json:"blocks"`
}

type EditorJsModelBlocks struct {
	Data EditorJsModelBlockText `json:"data"`
}

type EditorJsModelBlockText struct {
	Text string `json:"text"`
}
