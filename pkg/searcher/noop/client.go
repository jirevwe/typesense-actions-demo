package noop

import (
	"github.com/jirevwe/typesense-actions-demo/utils"
)

type NoopSearcher struct {
}

func NewNoopSearcher() *NoopSearcher {
	return &NoopSearcher{}
}

func (n *NoopSearcher) Search(collection string, filter *utils.SearchFilter) ([]string, utils.PaginationData, error) {
	return make([]string, 0), utils.PaginationData{}, nil
}

func (n *NoopSearcher) Index(collection string, document map[string]interface{}) error {
	return nil
}

func (n *NoopSearcher) Remove(collection string, filter *utils.SearchFilter) error {
	return nil
}
