package handlers_test

import (
	"fmt"
	"testing"

	"github.com/braydend/birthday-jamz/src/handlers"
)

func expectedButGot(expected, got string) string {
	return fmt.Sprintf("expected: %s. But got: %s", expected, got);
}

func TestHelloWorldHandler (t *testing.T) {
	result := handlers.HelloWorldHandler()

	if (result != "Hello, World") {
		t.Errorf("Handler returned wrong value");
	}
}

func TestHelloToHandler (t *testing.T) {
	expected:="Hello, Stephen!"
	result := handlers.HelloToHandler("Stephen")

	if (result != expected) {
		t.Errorf(expectedButGot(expected, result));
	}
}