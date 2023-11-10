package main

import (
	"fmt"
	"os"

	"github.com/YiumPotato/hypergofram/internal/routes"
)

func main() {
	// PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()

	r.Run(fmt.Sprintf(":%v", port))

}
