package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func newTestController(method, path string) (*beego.Controller, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	ctx := beegoCtx.NewContext()
	ctx.Reset(w, req)
	c := &beego.Controller{}
	c.Init(ctx, "TestController", method, nil)
	return c, w
}

func TestSendSuccessBuildsSuccessfulJSONResponse(t *testing.T) {
	c, w := newTestController(http.MethodGet, "/")
	SendSuccess(c, map[string]string{"message": "done"})

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, `"success":true`) {
		t.Fatalf("expected success:true in body, got %q", body)
	}
	if !strings.Contains(body, "done") {
		t.Fatalf("expected 'done' in body, got %q", body)
	}
}

func TestSendErrorBuildsErrorJSONResponse(t *testing.T) {
	c, w := newTestController(http.MethodGet, "/")
	SendError(c, http.StatusInternalServerError, "boom")

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, `"success":false`) {
		t.Fatalf("expected success:false, got %q", body)
	}
	if !strings.Contains(body, "boom") {
		t.Fatalf("expected 'boom' in body, got %q", body)
	}
}

func TestSendError400(t *testing.T) {
	c, w := newTestController(http.MethodPost, "/api/wishlist")
	SendError(c, http.StatusBadRequest, "invalid input")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestSendError404(t *testing.T) {
	c, w := newTestController(http.MethodGet, "/api/wishlist/missing")
	SendError(c, http.StatusNotFound, "not found")

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestJSONResponseSuccessStructure(t *testing.T) {
	r := JSONResponse{Success: true, Data: "hello"}
	if !r.Success {
		t.Error("expected Success=true")
	}
	if r.Data != "hello" {
		t.Errorf("expected Data=hello, got %v", r.Data)
	}
}

func TestJSONResponseErrorStructure(t *testing.T) {
	r := JSONResponse{Success: false, Message: "error occurred"}
	if r.Success {
		t.Error("expected Success=false")
	}
	if r.Message != "error occurred" {
		t.Errorf("expected message 'error occurred', got %s", r.Message)
	}
}

func TestSendSuccessWithSliceData(t *testing.T) {
	c, w := newTestController(http.MethodGet, "/api/countries")
	SendSuccess(c, []string{"France", "Japan"})

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "France") {
		t.Error("expected France in response")
	}
}

func TestSendSuccessWithNilData(t *testing.T) {
	c, w := newTestController(http.MethodGet, "/api/test")
	SendSuccess(c, nil)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"success":true`) {
		t.Error("expected success:true")
	}
}
