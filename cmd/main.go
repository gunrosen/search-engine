package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gunrosen/search-engine/internal/games"
	"github.com/gunrosen/search-engine/internal/guilds"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func init() {
	godotenv.Load()
}

func main() {
	os.Getenv("DB_HOST")
	searchCfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_URL")},
	}
	es, err := elasticsearch.NewClient(searchCfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	app := &cli.App{
		Name:  "Ingest search data",
		Usage: "GameFi.org Search",
		Commands: []*cli.Command{
			createIndexMappingCommand(),
			indexByCronJob(),
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
			log.Println("Done creating elastic mapping")
			return nil
		},
	}
}

func indexByCronJob() *cli.Command {
	return &cli.Command{
		Name:  "index-by-cron",
		Usage: "do index all by cron job.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "cron-job",
				Usage: "Cron job in '* * * * *' format. Option to run periodically. Default is not set.",
			},
		},
		Action: func(c *cli.Context) error {
			quit := make(chan os.Signal, 1)
			cronJob := c.String("cron-job")
			if cronJob != "" {
				log.Println("Running periodically with cron job:", cronJob)
				s := gocron.NewScheduler(time.UTC)
				s.Cron(cronJob).Do(actionIndexAll, c)
				s.StartBlocking()
			} else {
				err := actionIndexAll(c)
				if err != nil {
					log.Println(err)
				}
				quit <- os.Kill
				signal.Notify(quit, os.Interrupt)
				<-quit
			}
			return nil
		},
	}
}

func index() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "do index all.",
		Action: func(c *cli.Context) error {
			return actionIndexAll(c)
		},
	}
}

func indexGameCommand() *cli.Command {
	return &cli.Command{
		Name:  "index-game",
		Usage: "index game.",
		Action: func(c *cli.Context) error {
			return actionIndexGame(c)
		},
	}
}

func indexGuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "index-guild",
		Usage: "index guild.",
		Action: func(c *cli.Context) error {
			return actionIndexGuild(c)
		},
	}
}

func actionIndexAll(c *cli.Context) error {
	dt := time.Now()
	err := actionIndexGame(c)
	if err != nil {
		log.Println(err)
		return err
	}
	err = actionIndexGuild(c)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Done: last index %s", dt.Format("02-01-2006 15:04:05"))
	return nil
}

func actionIndexGame(c *cli.Context) error {
	es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
	if !ok {
		return errors.New("invalid Elasticsearch")
	}

	err := games.IngestData(es)
	if err != nil {
		return err
	}
	log.Println("Done index-game")
	return nil
}

func actionIndexGuild(c *cli.Context) error {
	es, ok := c.App.Metadata["db-search"].(*elasticsearch.Client)
	if !ok {
		return errors.New("invalid Elasticsearch")
	}

	err := guilds.IngestData(es)
	if err != nil {
		return err
	}
	log.Println("Done index-guild")
	return nil
}
