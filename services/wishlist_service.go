package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"travelsphere/models"
)

// In-memory store 
var (
	wishlistStore = make(map[string]models.WishlistItem)
	wishlistMu    sync.RWMutex
)

// generateID creates a simple unique ID without external packages
func generateID() string {
	return fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(9999))
}

// GetWishlist returns all wishlist items
func GetWishlist() []models.WishlistItem {
	wishlistMu.RLock()
	defer wishlistMu.RUnlock()

	items := make([]models.WishlistItem, 0, len(wishlistStore))
	for _, item := range wishlistStore {
		items = append(items, item)
	}
	return items
}


func AddToWishlist(countryName, note, status string) (*models.WishlistItem, error) {
	if countryName == "" {
		return nil, fmt.Errorf("country_name is required")
	}
	if status == "" {
		status = "Planned"
	}
	if status != "Planned" && status != "Visited" {
		return nil, fmt.Errorf("status must be Planned or Visited")
	}

	item := models.WishlistItem{
		ID:          generateID(),
		CountryName: countryName,
		Note:        note,
		Status:      status,
		CreatedAt:   time.Now(),
	}

	wishlistMu.Lock()
	wishlistStore[item.ID] = item
	wishlistMu.Unlock()

	return &item, nil
}


func UpdateWishlistItem(id, note, status string) (*models.WishlistItem, error) {
	if status != "Planned" && status != "Visited" {
		return nil, fmt.Errorf("status must be Planned or Visited")
	}

	wishlistMu.Lock()
	defer wishlistMu.Unlock()

	item, ok := wishlistStore[id]
	if !ok {
		return nil, fmt.Errorf("wishlist item not found")
	}
	item.Note = note
	item.Status = status
	wishlistStore[id] = item

	return &item, nil
}


func DeleteWishlistItem(id string) error {
	wishlistMu.Lock()
	defer wishlistMu.Unlock()

	if _, ok := wishlistStore[id]; !ok {
		return fmt.Errorf("wishlist item not found")
	}
	delete(wishlistStore, id)
	return nil
}
