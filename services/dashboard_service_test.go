package services

import (
	"testing"
	"travelsphere/models"
)

func resetForDashboard() {
	wishlistMu.Lock()
	defer wishlistMu.Unlock()
	wishlistStore = make(map[string]models.WishlistItem)
}

func TestGetDashboardSummaryEmpty(t *testing.T) {
	resetForDashboard()
	s := GetDashboardSummary()
	if s.Total != 0 || s.Planned != 0 || s.Visited != 0 {
		t.Errorf("expected all zeros, got total=%d planned=%d visited=%d", s.Total, s.Planned, s.Visited)
	}
	if s.Items == nil {
		t.Error("expected Items to be non-nil slice")
	}
}

func TestGetDashboardSummaryCountsPlanned(t *testing.T) {
	resetForDashboard()
	AddToWishlist("France", "", "Planned")
	AddToWishlist("Japan", "", "Planned")
	s := GetDashboardSummary()
	if s.Total != 2 {
		t.Errorf("expected total 2, got %d", s.Total)
	}
	if s.Planned != 2 {
		t.Errorf("expected planned 2, got %d", s.Planned)
	}
	if s.Visited != 0 {
		t.Errorf("expected visited 0, got %d", s.Visited)
	}
}

func TestGetDashboardSummaryCountsVisited(t *testing.T) {
	resetForDashboard()
	AddToWishlist("Italy", "Rome", "Visited")
	s := GetDashboardSummary()
	if s.Visited != 1 {
		t.Errorf("expected visited 1, got %d", s.Visited)
	}
}

func TestGetDashboardSummaryMixed(t *testing.T) {
	resetForDashboard()
	AddToWishlist("Germany", "", "Planned")
	AddToWishlist("Italy", "", "Visited")
	AddToWishlist("Spain", "", "Planned")
	s := GetDashboardSummary()
	if s.Total != 3 {
		t.Errorf("expected total 3, got %d", s.Total)
	}
	if s.Planned != 2 {
		t.Errorf("expected planned 2, got %d", s.Planned)
	}
	if s.Visited != 1 {
		t.Errorf("expected visited 1, got %d", s.Visited)
	}
}

func TestGetDashboardSummaryItemsContainData(t *testing.T) {
	resetForDashboard()
	AddToWishlist("Bangladesh", "Visit Dhaka", "Planned")
	s := GetDashboardSummary()
	if len(s.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(s.Items))
	}
	if s.Items[0].CountryName != "Bangladesh" {
		t.Errorf("expected Bangladesh, got %s", s.Items[0].CountryName)
	}
}
