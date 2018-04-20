package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/grafana/grafana/pkg/services/sqlstore/migrator"
)

var (
	repo             = "grafana/grafana"
	connectionString = ""
	apiToken         = ""
	database         = ""
	archiveUrl       = ""
	repoIds          = []int64{15111821}
)

func main() {
	flag.StringVar(&repo, "repo", "grafana/grafana", "name of the repo you want to process")
	flag.StringVar(&connectionString, "connectionString", "", "description")
	flag.StringVar(&database, "database", "", "description")
	flag.StringVar(&apiToken, "apiToken", "default?", "description")
	flag.StringVar(&archiveUrl, "archiveUrl", "default?", "description")
	flag.Parse()

	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	err = initDatabase(database, connectionString)
	if err != nil {
		log.Fatalf("migration failed. error: %v", err)
	}

	downloadEvents()
}