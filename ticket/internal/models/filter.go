package models

import (
	"net/url"

	"github.com/rodkevich/ts/ticket/pkg/filter"
)

type Filter struct {
	Base filter.Common

	Field1 string
	Field2 string
	Field3 string
}

func Filters(queries url.Values) *Filter {
	f := filter.New(queries)
	return &Filter{
		Base:   *f,
		Field1: queries.Get("1"),
		Field2: queries.Get("2"),
		Field3: queries.Get("3"),
	}
}
