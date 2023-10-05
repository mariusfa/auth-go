package routes

import (
	userController "auth/rest/user/controller"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

)

func TestLogin(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"username":"admin","password":"admin1"}`)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/user", requestBody)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}

	if (!strings.Contains(response.Body.String(), "token")) {
		t.Errorf("Response body is %v", response.Body.String())
	}
}

func TestProtectedPath(t *testing.T) {
	router := SetupRouter()
	token, err := userController.CreateToken(userController.SecretKey, map[string]interface{}{
		"username": "admin",
	})
	if (err != nil) {
		t.Errorf("Error creating token")
	}

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	
	// Add Authorization header
	request.Header.Add("Authorization", "Bearer "+token)

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Response code is %v", response.Code)
	}

	expected := "admin"
	if response.Body.String() != expected {
		t.Errorf("Response body is %v", response.Body.String())
	}
}
