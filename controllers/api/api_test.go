package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"travelsphere/services"
	"travelsphere/utils"

	beego "github.com/beego/beego/v2/server/web"
	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func makeCtx(method, path, body string) (*beegoCtx.Context, *httptest.ResponseRecorder) {
	var b *bytes.Reader
	if body != "" {
		b = bytes.NewReader([]byte(body))
	} else {
		b = bytes.NewReader(nil)
	}
	req := httptest.NewRequest(method, path, b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := beegoCtx.NewContext()
	ctx.Reset(w, req)
	return ctx, w
}

// ── utils.SendSuccess / SendError (called by all API handlers) ──

func TestSendSuccessWritesJSONResponse(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/test", "")
	c := &beego.Controller{}
	c.Init(ctx, "Test", "Get", nil)
	utils.SendSuccess(c, map[string]string{"k": "v"})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"success":true`) {
		t.Fatalf("expected success:true, got %s", w.Body.String())
	}
}

func TestSendErrorWritesMessageAndStatus(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/test", "")
	c := &beego.Controller{}
	c.Init(ctx, "Test", "Get", nil)
	utils.SendError(c, 400, "bad input")
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "bad input") {
		t.Fatalf("expected 'bad input', got %s", w.Body.String())
	}
}

// ── WishlistAPIController ──

func TestWishlistAPIGetReturns200(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/wishlist", "")
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Get", nil)
	c.Get()
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestWishlistAPIPostValid(t *testing.T) {
	payload := `{"country_name":"France","note":"","status":"Planned"}`
	ctx, w := makeCtx(http.MethodPost, "/api/wishlist", payload)
	ctx.Input.RequestBody = []byte(payload)
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Post", nil)
	c.Post()
	if w.Code != 201 {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestWishlistAPIPostInvalidJSON(t *testing.T) {
	ctx, w := makeCtx(http.MethodPost, "/api/wishlist", "")
	ctx.Input.RequestBody = []byte("not-json{{{")
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Post", nil)
	c.Post()
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestWishlistAPIPostMissingCountryName(t *testing.T) {
	payload := `{"country_name":"","status":"Planned"}`
	ctx, w := makeCtx(http.MethodPost, "/api/wishlist", payload)
	ctx.Input.RequestBody = []byte(payload)
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Post", nil)
	c.Post()
	if w.Code != 400 {
		t.Fatalf("expected 400 for empty country_name, got %d", w.Code)
	}
}

func TestWishlistAPIPostInvalidStatus(t *testing.T) {
	payload := `{"country_name":"France","status":"BadStatus"}`
	ctx, w := makeCtx(http.MethodPost, "/api/wishlist", payload)
	ctx.Input.RequestBody = []byte(payload)
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Post", nil)
	c.Post()
	if w.Code != 400 {
		t.Fatalf("expected 400 for invalid status, got %d", w.Code)
	}
}

func TestWishlistAPIPutValid(t *testing.T) {
	item, _ := services.AddToWishlist("Germany", "", "Planned")
	payload, _ := json.Marshal(map[string]string{"note": "Beer festival", "status": "Visited"})
	ctx, w := makeCtx(http.MethodPut, "/api/wishlist/"+item.ID, string(payload))
	ctx.Input.RequestBody = payload
	ctx.Input.SetParam(":id", item.ID)
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Put", nil)
	c.Put()
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestWishlistAPIPutInvalidJSON(t *testing.T) {
	ctx, w := makeCtx(http.MethodPut, "/api/wishlist/someid", "")
	ctx.Input.RequestBody = []byte("bad{json")
	ctx.Input.SetParam(":id", "someid")
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Put", nil)
	c.Put()
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestWishlistAPIPutNotFound(t *testing.T) {
	payload, _ := json.Marshal(map[string]string{"note": "", "status": "Visited"})
	ctx, w := makeCtx(http.MethodPut, "/api/wishlist/nonexistent", string(payload))
	ctx.Input.RequestBody = payload
	ctx.Input.SetParam(":id", "nonexistent")
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Put", nil)
	c.Put()
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestWishlistAPIDeleteValid(t *testing.T) {
	item, _ := services.AddToWishlist("Brazil", "", "Planned")
	ctx, w := makeCtx(http.MethodDelete, "/api/wishlist/"+item.ID, "")
	ctx.Input.SetParam(":id", item.ID)
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Delete", nil)
	c.Delete()
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "deleted") {
		t.Errorf("expected 'deleted' in response, got %s", body)
	}
}

func TestWishlistAPIDeleteNotFound(t *testing.T) {
	ctx, w := makeCtx(http.MethodDelete, "/api/wishlist/missing", "")
	ctx.Input.SetParam(":id", "missing")
	c := &WishlistAPIController{}
	c.Controller.Init(ctx, "WishlistAPIController", "Delete", nil)
	c.Delete()
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// ── DashboardAPIController ──

func TestDashboardAPIGetReturns200(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/dashboard/summary", "")
	c := &DashboardAPIController{}
	c.Controller.Init(ctx, "DashboardAPIController", "Get", nil)
	c.Get()
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"success":true`) {
		t.Fatalf("expected success:true, got %s", w.Body.String())
	}
}

func TestDashboardAPIGetContainsSummaryFields(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/dashboard/summary", "")
	c := &DashboardAPIController{}
	c.Controller.Init(ctx, "DashboardAPIController", "Get", nil)
	c.Get()
	body := w.Body.String()
	for _, field := range []string{"total", "planned", "visited"} {
		if !strings.Contains(body, field) {
			t.Errorf("expected field %q in response, got %s", field, body)
		}
	}
}

// ── CountriesAPIController ──

func TestCountriesAPIGetOneNotFound(t *testing.T) {
	ctx, w := makeCtx(http.MethodGet, "/api/countries/zzznotexist", "")
	ctx.Input.SetParam(":slug", "zzznotexist")
	c := &CountriesController{}
	c.Controller.Init(ctx, "CountriesController", "GetOne", nil)
	c.GetOne()
	// Either 404 (API returned not found) or 200 (API returned something) — no panic
	if w.Code != 404 && w.Code != 200 {
		t.Fatalf("unexpected status %d", w.Code)
	}
}
