package handlers_test

import (
	"reflect"
	"testing"

	"github.com/braydend/birthday-jamz/src/api"
	"github.com/braydend/birthday-jamz/src/handlers"
)

func TestPlaylistHandlerTakesValidDate (t *testing.T) {
	result, err := handlers.BuildPlaylistHandler("1989-01-19");

	if (reflect.DeepEqual(result, handlers.PlaylistResponse{Date: "1989-01-19", Songs: []api.Song{}})) {
		t.Logf(err.Error())
		t.Errorf("Handler returned wrong value");
	}
}

func TestPlaylistHandlerFailsOnWrongDate (t *testing.T) {
	result, err := handlers.BuildPlaylistHandler("1989-19-01");

	if (err == nil) {
		t.Logf("%v", result)
		t.Errorf("Handler expected error but none was provided");
	}
}