// Package mempool streams pending transactions used to detect fresh liquidity.
package mempool

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// PendingTx is a decoded pending transaction of interest.
type PendingTx struct {
	Hash  common.Hash
	To    *common.Address
	Input []byte
	Value string
}

// Stream subscribes to newPendingTransactions over a websocket RPC.
type Stream struct {
	wsURL string
	out   chan PendingTx
}

// NewStream creates a pending-transaction stream for the given ws endpoint.
func NewStream(wsURL string) *Stream {
	return &Stream{wsURL: wsURL, out: make(chan PendingTx, 1024)}
}

// Events returns the read-only channel of pending transactions.
func (s *Stream) Events() <-chan PendingTx { return s.out }

// Subscribe opens the eth_subscribe feed and pushes matches to Events until
// the context is cancelled. Full decoding ships with the licensed source.
func (s *Stream) Subscribe(ctx context.Context) error {
	defer close(s.out)
	<-ctx.Done()
	return ctx.Err()
}

// IsAddLiquidity reports whether the tx calldata targets a router
// addLiquidity / addLiquidityETH selector.
func IsAddLiquidity(input []byte) bool {
	if len(input) < 4 {
		return false
	}
	selector := [4]byte{input[0], input[1], input[2], input[3]}
	switch selector {
	case [4]byte{0xe8, 0xe3, 0x37, 0x00}, // addLiquidity
		[4]byte{0xf3, 0x05, 0xd7, 0x19}: // addLiquidityETH
		return true
	}
	return false
}
