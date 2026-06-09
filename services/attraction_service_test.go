package services

import (
	"os"
	"testing"
)

func TestGetAttractionsNoAPIKey(t *testing.T) {
	os.Unsetenv("OPENTRIPMAP_KEY")
	_, err := GetAttractions(23.7, 90.4, 5)
	if err == nil {
		t.Error("expected error when OPENTRIPMAP_KEY is missing")
	}
}

func TestGetAttractionsErrorMessage(t *testing.T) {
	os.Unsetenv("OPENTRIPMAP_KEY")
	_, err := GetAttractions(0, 0, 10)
	if err == nil {
		t.Error("expected error for missing key")
	}
	if err.Error() != "OpenTripMap key not set" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}
