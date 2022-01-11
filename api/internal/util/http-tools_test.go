package util

import (
	"net/url"
	"sort"
	"testing"
)

func TestConvertRowToFieldsArr(t *testing.T) {
	u, _ := url.Parse(
		"http://0.0.0.0:8000/ticket/list?fields[ticket]=some_field_to_be_returned,ticket_photos,ticket_description&fields[author]=name,id",
	)
	uv := u.Query()

	got, err := ConvertRowToFieldsArr(uv, "ticket")
	if err != nil {
		t.Errorf("ERROR: ConvertRowToFieldsArr: %v", err)

	}
	want := []string{"some_field_to_be_returned", "ticket_photos", "ticket_description"}

	sort.Strings(got)
	sort.Strings(want)

	for i, w := range want {
		g := got[i]
		if w != g {
			t.Errorf("got %v want %v", g, w)
		}
	}
}
