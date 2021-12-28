package filter

import (
	"net/url"
	"strconv"
)

const (
	// Query types
	queryParamFields = "fields"
	queryParamSearch = "search"
	// Paging settings
	queryParamPaging                = "paging"
	queryParamPage                  = "page"
	queryParamSize                  = "size"
	responseItemsDefaultPage        = 1
	responseItemsDefaultSizePerPage = 3
)

type Common struct {
	Extended bool   `json:"extended"`
	Search   bool   `json:"search"`
	Page     uint64 `json:"page"`
	Size     uint64 `json:"size"`
	Paging   bool   `json:"paging"`
}

func NewFromURL(queries url.Values) *Common {

	isExtended := has(queries, queryParamFields)
	isSearch := has(queries, queryParamSearch)
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
		Extended: isExtended,
		Search:   isSearch,
		Page:     uint64(page),
		Size:     uint64(sizePerPage),
		Paging:   paging,
	}
}

func has(queries url.Values, param string) bool {
	return queries.Get(param) != ""
}
