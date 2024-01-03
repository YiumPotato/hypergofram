package main

import (
	"fmt"
	"os"

	"github.com/YiumPotato/hypergofram/internal/db"
	"github.com/YiumPotato/hypergofram/internal/routes"
)

func main() {
	// PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()

	err := db.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.CloseDB()

	r.Run(fmt.Sprintf(":%v", port))

}
