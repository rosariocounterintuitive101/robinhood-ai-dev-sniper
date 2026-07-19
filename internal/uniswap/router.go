package uniswap

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// SwapParams describes a single exact-ETH-for-tokens swap.
type SwapParams struct {
	Router     common.Address
	Token      common.Address
	Recipient  common.Address
	AmountInWei *big.Int
	MinOutWei  *big.Int
	Deadline   int64
	GasPriceWei *big.Int
}

// Common mainnet routers. Override via config for other EVM chains.
var (
	UniswapV2Router = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	WETH            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

// Path returns the WETH->token swap path.
func (p SwapParams) Path() []common.Address {
	return []common.Address{WETH, p.Token}
}

// Valid reports whether the swap parameters are well-formed.
func (p SwapParams) Valid() bool {
	return p.AmountInWei != nil && p.AmountInWei.Sign() > 0 &&
		p.MinOutWei != nil && p.Token != (common.Address{})
}
