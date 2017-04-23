package helper

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "123456"
	hash, _ := HashPassword(password)

	if hash != password {
		t.Errorf("expected %v, got %v", password, hash)
	}
}
