// Package monitor tracks open positions and fans out live events over websockets.
package monitor

import (
	"sync"
	"time"
)

// State is the live tracking record for a monitored token.
type State struct {
	TokenAddress string
	CreatedAt    time.Time
	LastTxTime   time.Time
	TxCount      int64
	Profit       float64
}

// Message is a single websocket event pushed to subscribers.
type Message struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

// Client is a connected websocket subscriber.
type Client struct {
	ID   string
	Send chan Message
}

// Hub fans broadcast messages out to all clients grouped by token.
type Hub struct {
	mu        sync.RWMutex
	clients   map[string]map[*Client]bool
	broadcast chan Message
}

// NewHub constructs an empty hub ready to Run.
func NewHub() *Hub {
	return &Hub{
		clients:   make(map[string]map[*Client]bool),
		broadcast: make(chan Message, 256),
	}
}

// Subscribe registers a client under a token key.
func (h *Hub) Subscribe(token string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.clients[token] == nil {
		h.clients[token] = make(map[*Client]bool)
	}
	h.clients[token][c] = true
}

// Broadcast enqueues a message for delivery to all clients.
func (h *Hub) Broadcast(m Message) { h.broadcast <- m }

// Run delivers broadcast messages until the broadcast channel closes.
func (h *Hub) Run() {
	for msg := range h.broadcast {
		h.mu.RLock()
		for _, set := range h.clients {
			for c := range set {
				select {
				case c.Send <- msg:
				default:
				}
			}
		}
		h.mu.RUnlock()
	}
}
