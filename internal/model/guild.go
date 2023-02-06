package model

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
