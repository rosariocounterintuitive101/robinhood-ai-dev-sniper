// Package flashbots submits transaction bundles to the Flashbots relay
// for same-block, MEV-protected inclusion.
package flashbots

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/core/types"
)

// ErrPreviewBuild is returned by stubbed entry points in this public preview.
// The full implementation ships with the licensed source.
var ErrPreviewBuild = errors.New("not implemented in preview build — licensed source only")

// Bundle is an ordered set of signed transactions targeting a block.
type Bundle struct {
	Txs         types.Transactions
	BlockNumber uint64
	BribeWei    uint64
}

// Client talks to a Flashbots-compatible relay.
type Client struct {
	relayURL string
}

// NewClient returns a relay client for the given endpoint.
func NewClient(relayURL string) *Client { return &Client{relayURL: relayURL} }

// Send submits a signed bundle via eth_sendBundle and returns the bundle hash.
func (c *Client) Send(ctx context.Context, b *Bundle) (string, error) {
	if b == nil || len(b.Txs) == 0 {
		return "", errors.New("empty bundle")
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	return "", ErrPreviewBuild
}

// Simulate calls eth_callBundle to validate inclusion before submission.
func (c *Client) Simulate(ctx context.Context, b *Bundle) error {
	if b == nil || len(b.Txs) == 0 {
		return errors.New("empty bundle")
	}
	return ErrPreviewBuild
}
