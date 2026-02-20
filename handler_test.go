package go_auth_api_tests2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGetProduct(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		wantStatus int
	}{
		{
			name:       "200(valid)",
			url:        "/product?id=1",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "400(invalid/no id)",
			url:        "/product?id",
			method:     http.MethodGet,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "400(invalid/no int)",
			url:        "/product?id=abc",
			method:     http.MethodGet,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "405(invalid)",
			url:        "/product?id=1",
			method:     http.MethodPost,
			wantStatus: http.StatusOK,
		},
		{
			name:       "404(invalid)",
			url:        "/product?id=99",
			method:     http.MethodGet,
			wantStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rep := httptest.NewRecorder()
			HandlerGetProduct(rep, req)
			if rep.Code != tt.wantStatus {
				t.Errorf("want status %d, got %d", tt.wantStatus, rep.Code)
			}
		})
	}
}
