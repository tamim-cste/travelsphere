package services

import (
	"testing"
	"travelsphere/models"
)

func resetWishlistStore() {
	wishlistMu.Lock()
	defer wishlistMu.Unlock()
	wishlistStore = make(map[string]models.WishlistItem)
}

func TestAddToWishlistSuccess(t *testing.T) {
	resetWishlistStore()
	item, err := AddToWishlist("Bangladesh", "Visit Dhaka 2026", "Planned")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if item.CountryName != "Bangladesh" {
		t.Errorf("expected Bangladesh, got %s", item.CountryName)
	}
	if item.Status != "Planned" {
		t.Errorf("expected Planned, got %s", item.Status)
	}
	if item.ID == "" {
		t.Error("expected non-empty ID")
	}
	if item.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}
}

func TestAddToWishlistDefaultStatus(t *testing.T) {
	resetWishlistStore()
	item, err := AddToWishlist("France", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Status != "Planned" {
		t.Errorf("expected default status Planned, got %s", item.Status)
	}
}

func TestAddToWishlistInvalidStatus(t *testing.T) {
	resetWishlistStore()
	_, err := AddToWishlist("France", "", "InvalidStatus")
	if err == nil {
		t.Error("expected error for invalid status")
	}
}

func TestAddToWishlistEmptyName(t *testing.T) {
	resetWishlistStore()
	_, err := AddToWishlist("", "", "Planned")
	if err == nil {
		t.Error("expected error for empty country_name")
	}
}

func TestAddToWishlistVisitedStatus(t *testing.T) {
	resetWishlistStore()
	item, err := AddToWishlist("Japan", "Cherry blossom", "Visited")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Status != "Visited" {
		t.Errorf("expected Visited, got %s", item.Status)
	}
}

func TestGetWishlistReturnsItems(t *testing.T) {
	resetWishlistStore()
	AddToWishlist("Japan", "", "Visited")
	AddToWishlist("Brazil", "", "Planned")
	items := GetWishlist()
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestGetWishlistEmptyStore(t *testing.T) {
	resetWishlistStore()
	items := GetWishlist()
	if len(items) != 0 {
		t.Errorf("expected 0 items, got %d", len(items))
	}
}

func TestUpdateWishlistItemSuccess(t *testing.T) {
	resetWishlistStore()
	item, _ := AddToWishlist("Australia", "", "Planned")
	updated, err := UpdateWishlistItem(item.ID, "See Sydney Opera House", "Visited")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Status != "Visited" {
		t.Errorf("expected Visited, got %s", updated.Status)
	}
	if updated.Note != "See Sydney Opera House" {
		t.Errorf("expected note update, got %s", updated.Note)
	}
}

func TestUpdateWishlistItemNotFound(t *testing.T) {
	resetWishlistStore()
	_, err := UpdateWishlistItem("nonexistent-id", "note", "Planned")
	if err == nil {
		t.Error("expected error for nonexistent ID")
	}
}

func TestUpdateWishlistItemInvalidStatus(t *testing.T) {
	resetWishlistStore()
	item, _ := AddToWishlist("Germany", "", "Planned")
	_, err := UpdateWishlistItem(item.ID, "", "BadStatus")
	if err == nil {
		t.Error("expected error for invalid status")
	}
}

func TestDeleteWishlistItemSuccess(t *testing.T) {
	resetWishlistStore()
	item, _ := AddToWishlist("Brazil", "", "Planned")
	err := DeleteWishlistItem(item.ID)
	if err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}
	items := GetWishlist()
	if len(items) != 0 {
		t.Errorf("expected 0 items after delete, got %d", len(items))
	}
}

func TestDeleteWishlistItemNotFound(t *testing.T) {
	resetWishlistStore()
	err := DeleteWishlistItem("nonexistent")
	if err == nil {
		t.Error("expected error deleting non-existent item")
	}
}

func TestGenerateIDUniqueness(t *testing.T) {
	id1 := generateID()
	id2 := generateID()
	if id1 == id2 {
		t.Error("expected unique IDs")
	}
}
