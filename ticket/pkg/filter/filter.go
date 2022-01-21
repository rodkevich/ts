package filter

import (
	"net/url"
	"strconv"
)

const (
	queryParamSearch                = "search"
	queryParamPaging                = "paging"
	queryParamPage                  = "page"
	queryParamSize                  = "size"
	responseItemsDefaultPage        = 1
	responseItemsDefaultSizePerPage = 3
)

type Common struct {
	Search bool   `json:"search"`
	Page   uint64 `json:"page"`
	Size   uint64 `json:"size"`
	Paging bool   `json:"paging"`
}

func NewFromURL(queries url.Values) *Common {

	isSearch := has(queries, queryParamSearch)
	paging := has(queries, queryParamPaging)

	page, _ := strconv.Atoi(queries.Get(queryParamPage))
	if !has(queries, queryParamPage) {
		page = responseItemsDefaultPage
	}

	sizePerPage, _ := strconv.Atoi(queries.Get(queryParamSize))
	if !has(queries, queryParamSize) {
		sizePerPage = responseItemsDefaultSizePerPage
	}

	// offset calculation:
	page = (page - 1) * sizePerPage

	return &Common{
		Search: isSearch,
		Page:   uint64(page),
		Size:   uint64(sizePerPage),
		Paging: paging,
	}
}

func has(queries url.Values, param string) bool {
	return queries.Get(param) != ""
}
