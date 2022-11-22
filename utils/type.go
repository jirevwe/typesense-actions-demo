package utils

import (
	"fmt"
	"strings"
)

type Pageable struct {
	Page    int `json:"page" bson:"page"`
	PerPage int `json:"per_page" bson:"per_page"`
	Sort    int `json:"sort" bson:"sort"`
}

type SearchParams struct {
	CreatedAtStart int64 `json:"created_at_start" bson:"created_at_start"`
	CreatedAtEnd   int64 `json:"created_at_end" bson:"created_at_end"`
}

type FilterBy struct {
	AppID        string
	GroupID      string
	SourceID     string
	SearchParams SearchParams
}

func (f *FilterBy) String() *string {
	var s string
	filterByBuilder := new(strings.Builder)
	filterByBuilder.WriteString(fmt.Sprintf("group_id:=%s", f.GroupID))
	filterByBuilder.WriteString(fmt.Sprintf(" && created_at:[%d..%d]", f.SearchParams.CreatedAtStart, f.SearchParams.CreatedAtEnd))

	if len(f.AppID) > 0 {
		filterByBuilder.WriteString(fmt.Sprintf(" && app_id:=%s", f.AppID))
	}

	if len(f.SourceID) > 0 {
		filterByBuilder.WriteString(fmt.Sprintf(" && source_id:=%s", f.SourceID))
	}

	s = filterByBuilder.String()

	// we only return a pointer address here
	// because the typesense lib needs a string pointer
	return &s
}

type SearchFilter struct {
	Query    string
	FilterBy FilterBy
	Pageable Pageable
}

type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}
