package main

import (
	"encoding/json"
	"github/teacoder-team/golang-backend/config"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	config.ConnectDatabase(cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})

	log.Printf("🚀 Server is running at: %s\n", cfg.ApplicationURL)
	log.Printf("📄 Documentation is available at: %s%s\n", cfg.ApplicationURL, cfg.SwaggerPrefix)

	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.ApplicationPort), nil); err != nil {
		log.Fatalf("❌ Error starting server: %v\n", err)
	}
}
