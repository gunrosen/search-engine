package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gunrosen/search-engine/internal/games"
	"github.com/gunrosen/search-engine/internal/guilds"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func init() {
	godotenv.Load()
}

func main() {
	searchCfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_URL")},
	}
	es, err := elasticsearch.NewClient(searchCfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	app := &cli.App{
		Name:  "Search",
		Usage: "GameFi.org Search",
		Commands: []*cli.Command{
			createIndexMappingCommand(),
			index(),
			indexGameCommand(),
			indexGuildCommand(),
		},
		Metadata: map[string]any{
			"db-search": es,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createIndexMappingCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "create index mapping.",
		Action: func(c *cli.Context) error {
			es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
			if !ok {
				return errors.New("invalid Elasticsearch")
			}
			err := games.CreateMapping(es)
			if err != nil {
				return err
			}
			err = guilds.CreateMapping(es)
			if err != nil {
				return err
			}
			fmt.Println("Done creating elastic mapping")
			return nil
		},
	}
}

func index() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "do index all.",
		Action: func(c *cli.Context) error {
			es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
			if !ok {
				return errors.New("invalid Elasticsearch")
			}
			err := games.CreateMapping(es)
			if err != nil {
				return err
			}
			err = guilds.CreateMapping(es)
			if err != nil {
				return err
			}
			fmt.Println("Done index all.")
			return nil
		},
	}
}

func indexGameCommand() *cli.Command {
	return &cli.Command{
		Name:  "index-game",
		Usage: "index game.",
		Action: func(c *cli.Context) error {
			es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
			if !ok {
				return errors.New("invalid Elasticsearch")
			}
			err := games.IngestData(es)
			if err != nil {
				return err
			}
			fmt.Println("Done index-game")
			return nil
		},
	}
}

func indexGuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "index-guild",
		Usage: "index guild.",
		Action: func(c *cli.Context) error {
			es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
			if !ok {
				return errors.New("invalid Elasticsearch")
			}

			err := guilds.CreateMapping(es)
			if err != nil {
				return err
			}
			fmt.Println("Done index-guild")
			return nil
		},
	}
}
