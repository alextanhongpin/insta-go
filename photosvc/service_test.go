package photosvc

import (
	"testing"
)

type MockService struct{}

func (s MockService) All(request interface{}) ([]Photo, error) {
	var photos []Photo
	return photos, nil
}

func TestAll(t *testing.T) {
	actual := "hello"
	expected := "world"

	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
