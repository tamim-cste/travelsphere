package models

import (
	"reflect"
	"testing"
	"time"
)

func TestCountryAllFields(t *testing.T) {
	c := Country{
		Name: "Bangladesh", Capital: "Dhaka", Population: 170000000,
		Region: "Asia", Subregion: "Southern Asia",
		FlagURL: "https://flag.svg", Currency: "BDT", Languages: "Bengali",
		Slug: "bangladesh", Lat: 23.7, Lon: 90.4,
	}
	if c.Name != "Bangladesh" { t.Errorf("Name mismatch") }
	if c.Slug != "bangladesh" { t.Errorf("Slug mismatch") }
	if c.Lat != 23.7 { t.Errorf("Lat mismatch") }
	if c.Lon != 90.4 { t.Errorf("Lon mismatch") }
}

func TestCountryJSONTags(t *testing.T) {
	typ := reflect.TypeOf(Country{})
	checks := map[string]string{
		"Name": "name", "Capital": "capital", "Population": "population",
		"Region": "region", "FlagURL": "flag", "Slug": "slug",
		"Lat": "lat", "Lon": "lon",
	}
	for field, tag := range checks {
		f, ok := typ.FieldByName(field)
		if !ok { t.Errorf("field %s missing", field); continue }
		if got := f.Tag.Get("json"); got != tag {
			t.Errorf("field %s: expected json:%q got %q", field, tag, got)
		}
	}
}

func TestWishlistItemAllFields(t *testing.T) {
	now := time.Now()
	item := WishlistItem{
		ID: "abc", CountryName: "France", Note: "Eiffel",
		Status: "Planned", CreatedAt: now,
	}
	if item.ID != "abc" { t.Errorf("ID mismatch") }
	if item.CountryName != "France" { t.Errorf("CountryName mismatch") }
	if item.Note != "Eiffel" { t.Errorf("Note mismatch") }
	if item.Status != "Planned" { t.Errorf("Status mismatch") }
	if item.CreatedAt != now { t.Errorf("CreatedAt mismatch") }
}

func TestWishlistItemJSONTags(t *testing.T) {
	typ := reflect.TypeOf(WishlistItem{})
	checks := map[string]string{
		"ID": "id", "CountryName": "country_name",
		"Note": "note", "Status": "status", "CreatedAt": "created_at",
	}
	for field, tag := range checks {
		f, ok := typ.FieldByName(field)
		if !ok { t.Errorf("field %s missing", field); continue }
		if got := f.Tag.Get("json"); got != tag {
			t.Errorf("field %s: expected json:%q got %q", field, tag, got)
		}
	}
}

func TestWeatherAllFields(t *testing.T) {
	w := Weather{TempC: 30.0, Condition: "Sunny", Icon: "https://icon.png", Humidity: 80, WindKph: 10.5, City: "Dhaka"}
	if w.TempC != 30.0 { t.Errorf("TempC mismatch") }
	if w.Condition != "Sunny" { t.Errorf("Condition mismatch") }
	if w.City != "Dhaka" { t.Errorf("City mismatch") }
	if w.Humidity != 80 { t.Errorf("Humidity mismatch") }
	if w.WindKph != 10.5 { t.Errorf("WindKph mismatch") }
}

func TestAttractionAllFields(t *testing.T) {
	a := Attraction{Name: "Eiffel Tower", Kinds: "architecture,historic", XID: "Q243"}
	if a.Name != "Eiffel Tower" { t.Errorf("Name mismatch") }
	if a.Kinds != "architecture,historic" { t.Errorf("Kinds mismatch") }
	if a.XID != "Q243" { t.Errorf("XID mismatch") }
}

func TestUserField(t *testing.T) {
	u := User{Username: "beta"}
	if u.Username != "beta" { t.Errorf("Username mismatch") }
}

func TestUserJSONTag(t *testing.T) {
	typ := reflect.TypeOf(User{})
	f, ok := typ.FieldByName("Username")
	if !ok { t.Fatal("Username field missing") }
	if got := f.Tag.Get("json"); got != "username" {
		t.Errorf("expected json:username, got %q", got)
	}
}

func TestWishlistItemDefaultNote(t *testing.T) {
	item := WishlistItem{ID: "x", CountryName: "Japan", Status: "Visited", CreatedAt: time.Now()}
	if item.Note != "" {
		t.Errorf("expected empty default note, got %q", item.Note)
	}
}
