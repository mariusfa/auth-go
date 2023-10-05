package controller

import (
	"testing"
)

func TestValidToken(t *testing.T) {
	data := map[string]interface{}{
		"username": "admin",
	}
	token, err := CreateToken(SecretKey, data)
	if (err != nil) {
		t.Errorf("Error creating token")
	}
	user, err := ValidateToken(token)
	if (err != nil) {
		t.Errorf("Error validating token")
	}
	if user != "admin" {
		t.Errorf("User is not admin, got: %v", user)
	}

}

func TestInvalidUserInToken(t *testing.T) {
	data := map[string]interface{}{
		"username": "admin1",
	}
	token, err := CreateToken(SecretKey, data)
	if (err != nil) {
		t.Errorf("Error creating token")
	}
	_, err = ValidateToken(token)
	if (err == nil) {
		t.Errorf("Error validating token")
	}
}
