package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	token := "mytoken123"

	// Scenario 1: Success
	t.Run("Successfully get API key", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey "+token)
		got, err := GetAPIKey(headers)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := token
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	})

	//t.Run("Wrong Auth header with more than 2 strings", func(t *testing.T) {
	//	headers := http.Header{}
	//	headers.Set("Authorization", "ApiKey mytoken123 and something else")
	//	_, err := GetAPIKey(headers)
	//	if err == nil {
	//		t.Fatal("Supposed to get an error during getting API key with wrong header, but there is no error")
	//	}
	//})

	//t.Run("Wrong Auth header with empty key", func(t *testing.T) {
	//	headers := http.Header{}
	//	headers.Set("Authorization", "ApiKey ")
	//
	//	_, err := GetAPIKey(headers)
	//	if err == nil {
	//		t.Fatal("Supposed to get an error during getting API key, its empty")
	//	}
	//})

	t.Run("Empty Auth header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Fatal("wrong error message")
		}
	})

}
