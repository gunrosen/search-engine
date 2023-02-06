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

const IndexNameGame = "gamefi_games"

const Game = `
{
  "settings": {
	"index": {
		"max_ngram_diff" : "5"
	}
  },
  "mappings": {
    "properties": {
      "description": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "excerpt": {
        "type": "text"
      },
      "features_text": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "guideline": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "introduction": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "name": {
        "type": "text"
      },
      "play_mode": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "play_to_earn_model": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "roadmap_text": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      },
      "slug": {
        "type": "keyword"
      },
      "white_paper": {
        "properties": {
          "blocks": {
            "type": "nested",
            "properties": {
              "data": {
                "properties": {
                  "text": {
                    "type": "text"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
`
