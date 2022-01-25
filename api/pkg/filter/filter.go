package filter

import (
	"net/url"
	"strconv"
)

const (
	// Query types
	queryParamSearch = "search"
	// Paging settings
	queryParamPaging                = "paging"
	queryParamReversed              = "reversed"
	queryParamPage                  = "page"
	queryParamSize                  = "page_size"
	responseItemsDefaultPage        = 1
	responseItemsDefaultSizePerPage = 10
)

// Common struct to be used in composition of filters
type Common struct {
	Search   bool   `json:"search"`
	Reversed bool   `json:"reversed"`
	Page     uint64 `json:"page"`
	Size     uint64 `json:"size"`
	Paging   bool   `json:"paging"`
}

// NewFromURL create a new filter
func NewFromURL(queries url.Values) *Common {

	isSearch := has(queries, queryParamSearch)
	isReversed, _ := strconv.ParseBool(queries.Get(queryParamReversed))
	page, _ := strconv.Atoi(queries.Get(queryParamPage))
	sizePerPage, _ := strconv.Atoi(queries.Get(queryParamSize))
	paging, _ := strconv.ParseBool(queries.Get(queryParamPaging))

	if !has(queries, queryParamSize) {
		sizePerPage = responseItemsDefaultSizePerPage
	}

	if !has(queries, queryParamPage) {
		page = responseItemsDefaultPage
	}
	// offset calculation:
	page = (page - 1) * sizePerPage

	return &Common{
		Search:   isSearch,
		Page:     uint64(page),
		Size:     uint64(sizePerPage),
		Paging:   paging,
		Reversed: isReversed,
	}
}

func has(queries url.Values, param string) bool {
	return queries.Get(param) != ""
}
