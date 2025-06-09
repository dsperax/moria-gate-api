package main

import (
	"log"
	"moria-gate-api/internal/infra/http/router"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("[INFO] Starting server on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
