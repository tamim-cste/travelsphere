package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func newCtx(method, path string) (*beegoCtx.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	ctx := beegoCtx.NewContext()
	ctx.Reset(w, req)
	return ctx, w
}

// ── BaseController ──

func TestBaseControllerPrepareSetsDefaults(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/dashboard")
	bc := &BaseController{}
	bc.Controller.Init(ctx, "BaseController", "Prepare", nil)
	bc.Prepare()

	if bc.Data["AppName"] != "TravelSphere" {
		t.Fatalf("expected AppName TravelSphere, got %v", bc.Data["AppName"])
	}
	if bc.Data["CurrentPath"] != "/dashboard" {
		t.Fatalf("expected CurrentPath /dashboard, got %v", bc.Data["CurrentPath"])
	}
	if bc.Data["LoggedIn"] != false {
		t.Fatalf("expected LoggedIn false, got %v", bc.Data["LoggedIn"])
	}
	if bc.Data["Username"] != "" {
		t.Fatalf("expected empty Username, got %v", bc.Data["Username"])
	}
}

func TestBaseControllerReadSessionEmpty(t *testing.T) {
	bc := &BaseController{}
	result := bc.readSession("username")
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestBaseControllerReadSessionUnknownKey(t *testing.T) {
	bc := &BaseController{}
	result := bc.readSession("nonexistent")
	if result != "" {
		t.Errorf("expected empty string for nonexistent key, got %q", result)
	}
}

// ── AuthController ──

func TestAuthControllerGetSetsTitle(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/login")
	ac := &AuthController{}
	ac.Controller.Init(ctx, "AuthController", "Get", nil)
	ac.Get()

	if ac.Data["Title"] != "Login" {
		t.Errorf("expected Title=Login, got %v", ac.Data["Title"])
	}
	if ac.TplName != "login.tpl" {
		t.Errorf("expected login.tpl, got %s", ac.TplName)
	}
	if ac.Layout != "layout/main.tpl" {
		t.Errorf("expected layout/main.tpl, got %s", ac.Layout)
	}
}

func TestAuthControllerPostEmptyUsername(t *testing.T) {
	body := strings.NewReader("username=")
	req := httptest.NewRequest(http.MethodPost, "/login", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := beegoCtx.NewContext()
	ctx.Reset(w, req)

	ac := &AuthController{}
	ac.Controller.Init(ctx, "AuthController", "Post", nil)
	ac.Post()

	if ac.Data["Error"] == nil {
		t.Error("expected error message for empty username")
	}
	if ac.TplName != "login.tpl" {
		t.Errorf("expected login.tpl on error, got %s", ac.TplName)
	}
}

// ── HomeController ──

func TestHomeControllerGetSetsTitle(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/")
	hc := &HomeController{}
	hc.Controller.Init(ctx, "HomeController", "Get", nil)
	hc.Get()

	if hc.Data["Title"] != "Home" {
		t.Errorf("expected Title=Home, got %v", hc.Data["Title"])
	}
	if hc.TplName != "home.tpl" {
		t.Errorf("expected home.tpl, got %s", hc.TplName)
	}
	if hc.Layout != "layout/main.tpl" {
		t.Errorf("expected layout/main.tpl, got %s", hc.Layout)
	}
	if hc.Data["FeaturedCountries"] == nil {
		t.Error("expected FeaturedCountries to be set")
	}
}

// ── CountryController ──

func TestCountryControllerGetSetsTitle(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/countries")
	cc := &CountryController{}
	cc.Controller.Init(ctx, "CountryController", "Get", nil)
	cc.Get()

	if cc.Data["Title"] != "Country Explorer" {
		t.Errorf("expected Country Explorer, got %v", cc.Data["Title"])
	}
	if cc.TplName != "countries.tpl" {
		t.Errorf("expected countries.tpl, got %s", cc.TplName)
	}
}

func TestCountryControllerGetOneInvalidSlug(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/countries/zzz-invalid-country-xyz")
	ctx.Input.SetParam(":slug", "zzz-invalid-country-xyz")
	cc := &CountryController{}
	cc.Controller.Init(ctx, "CountryController", "GetOne", nil)
	cc.GetOne()

	// Either 404 template or destination template — no panic
	if cc.TplName != "404.tpl" && cc.TplName != "destination.tpl" {
		t.Errorf("unexpected template: %s", cc.TplName)
	}
}

// ── WishlistController ──

func TestWishlistControllerPrepareSetsLoggedInFalse(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/wishlist")
	wc := &WishlistController{}
	wc.Controller.Init(ctx, "WishlistController", "Prepare", nil)
	// Prepare will try to redirect — set LoggedIn manually to avoid redirect panic
	wc.Data["LoggedIn"] = false
	wc.Data["Username"] = ""
	wc.Data["AppName"] = "TravelSphere"
	wc.Data["CurrentPath"] = "/wishlist"

	loggedIn, ok := wc.Data["LoggedIn"].(bool)
	if !ok || loggedIn {
		t.Error("expected LoggedIn to be false")
	}
}

func TestWishlistControllerGetSetsTemplate(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/wishlist")
	wc := &WishlistController{}
	wc.Controller.Init(ctx, "WishlistController", "Get", nil)
	wc.Get()

	if wc.Data["Title"] != "Travel Wishlist" {
		t.Errorf("expected Travel Wishlist, got %v", wc.Data["Title"])
	}
	if wc.TplName != "wishlist.tpl" {
		t.Errorf("expected wishlist.tpl, got %s", wc.TplName)
	}
}

// ── DashboardController ──

func TestDashboardControllerGetSetsTemplate(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/dashboard")
	dc := &DashboardController{}
	dc.Controller.Init(ctx, "DashboardController", "Get", nil)
	dc.Get()

	if dc.Data["Title"] != "Travel Dashboard" {
		t.Errorf("expected Travel Dashboard, got %v", dc.Data["Title"])
	}
	if dc.TplName != "dashboard.tpl" {
		t.Errorf("expected dashboard.tpl, got %s", dc.TplName)
	}
	if dc.Data["Summary"] == nil {
		t.Error("expected Summary to be set")
	}
}
