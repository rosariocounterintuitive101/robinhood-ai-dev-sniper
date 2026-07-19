// Package uniswap implements constant-product AMM math and pair inspection.
package uniswap

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// PairReserves is a snapshot of a Uniswap V2 style pair.
type PairReserves struct {
	Reserve0, Reserve1 *big.Int
	Token0, Token1     common.Address
	BlockTimestampLast uint32
	TaxBps             uint16
	LiquidityLocked    bool
}

const (
	feeNumerator   = 997
	feeDenominator = 1000
)

// GetAmountOut applies the Uniswap V2 constant-product formula with the 0.3% fee.
func GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	if amountIn.Sign() <= 0 || reserveIn.Sign() <= 0 || reserveOut.Sign() <= 0 {
		return big.NewInt(0)
	}
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(feeNumerator))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(
		new(big.Int).Mul(reserveIn, big.NewInt(feeDenominator)),
		amountInWithFee,
	)
	return new(big.Int).Div(numerator, denominator)
}

// MinOut applies slippage tolerance (in bps) to an expected output amount.
func MinOut(expected *big.Int, slippageBps uint16) *big.Int {
	keep := big.NewInt(int64(10000 - slippageBps))
	out := new(big.Int).Mul(expected, keep)
	return out.Div(out, big.NewInt(10000))
}

// SpotPrice returns reserveOut/reserveIn scaled by 1e18.
func SpotPrice(reserveIn, reserveOut *big.Int) *big.Int {
	if reserveIn.Sign() == 0 {
		return big.NewInt(0)
	}
	scaled := new(big.Int).Mul(reserveOut, big.NewInt(1e18))
	return scaled.Div(scaled, reserveIn)
}
