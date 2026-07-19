package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/0xNikoDev/robinhood-ai-dev-sniper/internal/types"
)

func (s *Server) routes() {
	s.mux.HandleFunc("POST /api/snipe/fire", s.handleSnipeFire)
	s.mux.HandleFunc("POST /api/snipe/bundle-fire", s.handleBundleFire)
	s.mux.HandleFunc("POST /api/sell/bundle", s.handleSellBundle)
	s.mux.HandleFunc("POST /api/token/create", s.handleTokenCreate)
	s.mux.HandleFunc("GET /api/token/", s.handleTokenInfo)
}

// withAuth enforces a Bearer license key on every request.
func (s *Server) withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !s.validLicense(token) {
			writeJSON(w, http.StatusUnauthorized, map[string]any{
				"success": false, "error": "invalid license key",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) validLicense(key string) bool {
	if key == "" {
		return false
	}
	for _, k := range s.cfg.LicenseKeys {
		if strings.TrimSpace(k) == key {
			return true
		}
	}
	return false
}

func (s *Server) handleSnipeFire(w http.ResponseWriter, r *http.Request) {
	var req types.SnipeFireRequest
	if err := decode(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, errResp(err.Error()))
		return
	}
	resp, _ := s.snipes.Fire(r.Context(), req)
	writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleBundleFire(w http.ResponseWriter, r *http.Request) {
	var req types.BundleSnipeFireRequest
	if err := decode(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, errResp(err.Error()))
		return
	}
	resp, _ := s.snipes.BundleFire(r.Context(), req)
	writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleSellBundle(w http.ResponseWriter, r *http.Request) {
	var req types.SellBundleRequest
	if err := decode(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, errResp(err.Error()))
		return
	}
	resp, _ := s.snipes.SellBundle(r.Context(), req)
	writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleTokenCreate(w http.ResponseWriter, r *http.Request) {
	var req types.TokenCreateRequest
	if err := decode(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, errResp(err.Error()))
		return
	}
	resp, _ := s.tokens.Deploy(r.Context(), req)
	writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleTokenInfo(w http.ResponseWriter, r *http.Request) {
	address := strings.TrimPrefix(r.URL.Path, "/api/token/")
	if address == "" {
		writeJSON(w, http.StatusBadRequest, errResp("token address required"))
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"tokenAddress": address, "error": "preview build"})
}

func decode(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func errResp(msg string) map[string]any {
	return map[string]any{"success": false, "error": msg}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
