package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T){
headers := make(http.Header)
	headers.Set("Authorization", "ApiKey abc123")
	
	// Call the function
	apiKey, err := GetAPIKey(headers)
	
	// Check for success
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}
	
	// Check correct API key was extracted
	expectedKey := "abc123"
	if apiKey != expectedKey {
		t.Errorf("Expected API key '%s', but got '%s'", expectedKey, apiKey)
	}

}
