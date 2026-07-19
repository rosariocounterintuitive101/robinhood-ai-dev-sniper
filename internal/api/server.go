// Package api exposes the REST + websocket control surface.
package api

import (
	"net/http"

	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/config"
	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/flashbots"
	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/monitor"
	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/sniper"
	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/token"
)

// Server bundles the router, engine managers and shared config.
type Server struct {
	cfg     *config.Config
	mux     *http.ServeMux
	snipes  *sniper.Manager
	tokens  *token.Creator
	hub     *monitor.Hub
}

// NewServer constructs a fully wired API server.
func NewServer(cfg *config.Config) *Server {
	fb := flashbots.NewClient(cfg.FlashbotsURL)
	hub := monitor.NewHub()
	go hub.Run()

	s := &Server{
		cfg:    cfg,
		mux:    http.NewServeMux(),
		snipes: sniper.NewManager(cfg.RPCURL, fb),
		tokens: token.NewCreator(cfg.RPCURL),
		hub:    hub,
	}
	s.routes()
	return s
}

// Listen serves HTTP with the license-auth middleware applied.
func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.withAuth(s.mux))
}
