package util

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

// ConvertStringArrToUUID convert string slice to uuid
func ConvertStringArrToUUID(ids []string) ([]uuid.UUID, error) {
	uids := make([]uuid.UUID, 0, len(ids))
	for _, id := range ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		uids = append(uids, uid)
	}

	return uids, nil
}

func FieldsFromURL(values url.Values, name string) ([]string, error) {
	if name == "" || values == nil {

		return nil, nil
	}

	key := fmt.Sprintf("fields[%s]", name)
	raw, ok := values[key]
	if ok {
		substrings := strings.Split(raw[0], ",")
		var rtn []string
		rtn = append(rtn, substrings...)

		return rtn, nil
	}

	return nil, nil
}
