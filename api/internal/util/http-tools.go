package util

import (
	"net/url"

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

func ConvertRowToFieldsArr(values url.Values, name string) ([]string, error) {
	var rtn []string
	// []string{"some_field_to_be_returned", "ticket_photos", "ticket_description"}

	return rtn, nil
}
