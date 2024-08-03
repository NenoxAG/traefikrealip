package traefikrealip_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NenoxAG/traefikrealip"
)

func TestPlugin(t *testing.T) {
	cfg := traefikrealip.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikrealip.New(ctx, next, cfg, "traefikrealip")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Forwarded-For", "10.1.0.0, 10.2.0.0, 10.3.0.0")

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Real-Ip", "10.1.0.0")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
