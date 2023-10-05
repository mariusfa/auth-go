package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router := SetupRouter()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}
	expected := "ok"
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}