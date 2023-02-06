package model

type Game struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Excerpt         string `json:"excerpt"`
	Description     string `json:"description"`
	FeaturesText    string `json:"features_text"`
	RoadmapText     string `json:"roadmap_text"`
	WhitePaper      string `json:"white_paper"`
	Introduction    string `json:"introduction"`
	PlayMode        string `json:"play_mode"`
	PlayToEarnModel string `json:"play_to_earn_model"`
	Guideline       string `json:"guideline"`
}

// -------------- ElasticSearch Model ------------
const (
	IndexNameGame   = "gamefi_games"
	GameMappingType = `
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
)
