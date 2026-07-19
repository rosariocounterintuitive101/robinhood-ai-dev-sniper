// Command robinhood starts the RobinHood AI Dev Sniper API server.
package main

import (
	"log"
	"os"

	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/api"
	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	srv := api.NewServer(cfg)
	log.Printf("RobinHood AI Dev Sniper — listening on %s (chain %d)", addr, cfg.ChainID)
	if err := srv.Listen(addr); err != nil {
		log.Fatalf("server: %v", err)
	}
}
