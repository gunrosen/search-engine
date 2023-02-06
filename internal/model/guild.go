package model

type Guild struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Staffs        []Staff `json:"staffs"`
	Regions       string  `json:"regions"`
	MemberRegions string  `json:"member_regions"`
	Languages     string  `json:"languages"`
	Difference    string  `json:"difference"`
	Introduction  string  `json:"introduction"`
}

type Staff struct {
	Name string `json:"name"`
}

// -------------- ElasticSearch Model ------------
const IndexNameGuild = "gamefi_guilds"

const GuildMappingType = `
{
  "mappings": {
    "properties": {
      "staffs": {
          "type": "nested",
            "properties": {
 				"name": {
                    "type": "text"
				}
            }
      },
      "regions": {
        "type": "text"
      },
      "member_regions": {
        "type": "text"
      },
      "introduction": {
        "type": "text"
      },
      "languages": {
        "type": "text"
      },
      "difference": {
        "type": "text"
      },
      "name": {
        "type": "text"
      },
      "slug": {
        "type": "keyword"
      }
    }
  }
}
`
