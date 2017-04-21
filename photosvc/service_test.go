package photosvc

import (
	"testing"
)

type MockService struct{}

func (s MockService) All() ([]Photo, error) {
	var photos []Photo
	return photos, nil
}

func (s MockService) One(request interface{}) (Photo, error) {
	var photo Photo
	return photo, nil
}

func TestAll(t *testing.T) {
	var mockService MockService
	var photos []Photo
	actual, _ := mockService.All()
	expected := photos

	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
