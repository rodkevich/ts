package util

import (
	"net/url"
	"sort"
	"testing"
)

var (
	URL, _ = url.Parse(
		"http://0.0.0.0:8000/ticket/list?fields[ticket]=some_field_to_be_returned,ticket_photos,ticket_description&fields[author]=name,id",
	)
	urlValues = URL.Query()
)

func TestConvertRowToFieldsArr(t *testing.T) {
	got, err := FieldsFromURL(urlValues, "ticket")
	if err != nil {
		t.Errorf("ERROR: FieldsFromURL: %v", err)

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

func BenchmarkFieldsFromURL(b *testing.B) {
	// b.SetBytes(123)
	for i := 0; i < b.N; i++ {
		FieldsFromURL(urlValues, "ticket")
	}
}
