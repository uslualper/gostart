package client

import (
	"testing"

	"go-start/utils/http"
)

func TestClient(t *testing.T) {

	client := http.Client{}
	client.SetTimeout(1)
	client.Init("https://www.google.com")
	_, status := client.Get()

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

}
