package searcher

import (
	"github.com/frain-dev/convoy/config"
	"github.com/jirevwe/typesense-actions-demo/pkg/searcher/noop"
	"github.com/jirevwe/typesense-actions-demo/pkg/searcher/typesense"
	"github.com/jirevwe/typesense-actions-demo/utils"
)

type Searcher interface {
	// Search retrieves documents from the typesense collection based on the search filters
	Search(collection string, filter *utils.SearchFilter) ([]string, utils.PaginationData, error)

	// Index upserts the collection and indexes documents in the typesense collection,
	// each document must have the id, uid, created_at and updated_at fields
	Index(collection string, document map[string]interface{}) error

	// Remove removes documents from the typesense collection based on the search filters
	Remove(collection string, filter *utils.SearchFilter) error
}

func NewSearchClient(c config.Configuration) (Searcher, error) {
	if c.Search.Type == config.SearchProvider("typesense") {
		client, err := typesense.NewTypesenseClient(c.Search.Typesense.Host, c.Search.Typesense.ApiKey)
		return client, err
	}

	return noop.NewNoopSearcher(), nil
}
