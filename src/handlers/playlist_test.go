package handlers_test

import (
	"reflect"
	"testing"

	"github.com/braydend/birthday-jamz/src/api"
	"github.com/braydend/birthday-jamz/src/handlers"
)

func TestPlaylistHandlerTakesValidDate (t *testing.T) {
	result, _ := handlers.BuildPlaylistHandler("2020-01-19");
	expected:=handlers.PlaylistResponse{Date: "2020-01-19", Songs: map[string]api.Song{
		"2020-01-19": {Title: "Old Town Road", Artist:  "Lil Nas X Featuring Billy Ray Cyrus"},
		"2021-01-19": {Title: "Old Town Road", Artist:  "Lil Nas X Featuring Billy Ray Cyrus"},
		"2022-01-19": {Title: "Old Town Road", Artist:  "Lil Nas X Featuring Billy Ray Cyrus"},
		"2023-01-19": {Title: "Old Town Road", Artist:  "Lil Nas X Featuring Billy Ray Cyrus"},
		"2024-01-19": {Title: "Old Town Road", Artist:  "Lil Nas X Featuring Billy Ray Cyrus"},
	}}

	if (!reflect.DeepEqual(result, expected)) {
		t.Errorf("Handler returned wrong value.\nExpected: %v\nGot: %v", expected, result);
	}
}

func TestPlaylistHandlerFailsOnWrongDate (t *testing.T) {
	result, err := handlers.BuildPlaylistHandler("1989-19-01");

	if (err == nil) {
		t.Logf("%v", result)
		t.Errorf("Handler expected error but none was provided");
	}
}