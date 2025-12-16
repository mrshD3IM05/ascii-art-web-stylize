package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/server/handlers"
)

func main() {
	handlers.Handlers()

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server failed:", err)
	}
}
