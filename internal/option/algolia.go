package option

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/urfave/cli"
	"log"
	"os"
)

const apiID = "WT16MPP7YO"
const indexName = "ta2mo.github.io"

// GenerateAlgoliaRecord
func GenerateAlgoliaRecord(_ *cli.Context) error {
	posts, err := MappingFromFiles(postDir)
	if err != nil {
		return err
	}

	apiKey := os.Getenv("algolia_api_key")

	client := search.NewClient(apiID, apiKey)
	index := client.InitIndex(indexName)
	if _, err := index.SaveObjects(posts); err != nil {
		log.Print("[ERROR]Save object error. %v", err.Error())
		return err
	}

	log.Print("[INFO]Success save object")

	return nil
}
