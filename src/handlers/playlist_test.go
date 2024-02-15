package handlers_test

import (
	"testing"

	"github.com/braydend/birthday-jamz/src/handlers"
)

func TestPlaylistHandlerTakesValidDate (t *testing.T) {
	result, err := handlers.BuildPlaylistHandler("1989-01-19");

	if (result != "I would build a date starting on 1989-01-19") {
		t.Logf(err.Error())
		t.Errorf("Handler returned wrong value");
	}
}

func TestPlaylistHandlerFailsOnWrongDate (t *testing.T) {
	result, err := handlers.BuildPlaylistHandler("1989-19-01");

	if (err == nil) {
		t.Logf(result)
		t.Errorf("Handler expected error but none was provided");
	}
}