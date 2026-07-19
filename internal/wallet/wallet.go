// Package wallet manages EVM signing keys used for coordinated snipes.
package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet wraps a secp256k1 key and its derived address.
type Wallet struct {
	Address    common.Address
	privateKey *ecdsa.PrivateKey
}

// FromPrivateKey builds a Wallet from a hex-encoded private key.
func FromPrivateKey(hexKey string) (*Wallet, error) {
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	return &Wallet{
		Address:    crypto.PubkeyToAddress(pk.PublicKey),
		privateKey: pk,
	}, nil
}

// Key returns the underlying private key for transaction signing.
func (w *Wallet) Key() *ecdsa.PrivateKey { return w.privateKey }

// Set is an ordered collection of wallets rotated through during snipes.
type Set struct {
	wallets []*Wallet
}

// LoadSet parses a slice of hex private keys into a wallet Set.
func LoadSet(hexKeys []string) (*Set, error) {
	s := &Set{}
	for i, k := range hexKeys {
		w, err := FromPrivateKey(k)
		if err != nil {
			return nil, fmt.Errorf("wallet %d: %w", i, err)
		}
		s.wallets = append(s.wallets, w)
	}
	return s, nil
}

// All returns every wallet in the set.
func (s *Set) All() []*Wallet { return s.wallets }

// Len reports the number of wallets in the set.
func (s *Set) Len() int { return len(s.wallets) }
