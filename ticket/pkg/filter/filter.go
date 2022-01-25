package filter

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
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

	// var cursor string
	// if !has(queries, "cur") {
	// 	createdCursor, paymentID, errCsr := decodeCursor(params.Cursor)
	// 	if errCsr != nil {
	// 		err = errors.New("invalid-cursor")
	// 		return
	// 	}
	// 	queryBuilder = queryBuilder.Where(sq.LtOrEq{
	// 		"created_time": createdCursor,
	// 	})
	// 	queryBuilder = queryBuilder.Where(sq.Lt{
	// 		"id": paymentID,
	// 	})
	// }

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

// func DecodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
func DecodeCursor(encodedCursor string) (res string, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")

	println("\n DecodeCursor -- : ", arrStr[0], arrStr[1])

	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	// res, err = time.Parse(time.RFC3339Nano, arrStr[0])
	// if err != nil {
	// 	return
	// }
	// println("\n DecodeCursor -- : ", res.String())

	res = arrStr[0]
	uuid = arrStr[1]
	return
}

func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
