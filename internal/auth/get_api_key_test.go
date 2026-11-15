
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 123456")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if apiKey != "123456" {
		t.Errorf("expected 123456, got %s", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) 
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer 123456")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
