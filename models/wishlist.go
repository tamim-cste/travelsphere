package models

import "time"


type WishlistItem struct {
	ID          string    `json:"id"`
	CountryName string    `json:"country_name"`
	Note        string    `json:"note"`
	Status      string    `json:"status"`     // "Planned" or "Visited"
	CreatedAt   time.Time `json:"created_at"`
}
