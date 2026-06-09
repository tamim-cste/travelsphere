package services

import "travelsphere/models"


type DashboardSummary struct {
	Total   int                  `json:"total"`
	Planned int                  `json:"planned"`
	Visited int                  `json:"visited"`
	Items   []models.WishlistItem `json:"items"`
}


func GetDashboardSummary() DashboardSummary {
	items := GetWishlist()
	summary := DashboardSummary{Items: items}
	for _, item := range items {
		summary.Total++
		if item.Status == "Planned" {
			summary.Planned++
		} else if item.Status == "Visited" {
			summary.Visited++
		}
	}
	return summary
}
