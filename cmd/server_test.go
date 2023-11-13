// FILEPATH: /home/gg/workspace/hypergofram/cmd/server_test.go
package main

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/YiumPotato/hypergofram/internal/routes"
)

func TestMain(t *testing.T) {
	// Set the PORT environment variable
	os.Setenv("PORT", "8081")

	// Run the main function in a goroutine as it blocks
	go main()

	// Give the server some time to start
	time.Sleep(time.Second)

	// Test if the server is running
	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", resp.StatusCode)
	}
}

func TestSetupRouter(t *testing.T) {
	r := routes.SetupRouter()

	if r == nil {
		t.Fatalf("Expected a router, got nil")
	}
}
