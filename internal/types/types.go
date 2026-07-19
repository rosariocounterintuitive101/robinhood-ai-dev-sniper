// Package types holds the shared request/response payloads for the REST API.
package types

// SnipeFireRequest is the payload for POST /api/snipe/fire.
type SnipeFireRequest struct {
	TokenAddress string   `json:"tokenAddress"`
	Wallets      []string `json:"wallets"`
	BuyPercent   float64  `json:"buyPercent"`
	SlippageBps  uint16   `json:"slippageBps"`
	MaxGasGwei   uint64   `json:"maxGasGwei"`
}

// BundleSnipeFireRequest is the payload for POST /api/snipe/bundle-fire.
type BundleSnipeFireRequest struct {
	TokenAddress       string   `json:"tokenAddress"`
	Wallets            []string `json:"wallets"`
	BribeETH           float64  `json:"bribeETH"`
	RpcFallbackDelayMs int64    `json:"rpcFallbackDelayMs"`
}

// VolumeStartRequest is the payload for POST /api/volume/start.
type VolumeStartRequest struct {
	TokenAddress string   `json:"tokenAddress"`
	Wallets      []string `json:"wallets"`
	Cycles       int      `json:"cycles"`
	DelayMs      int64    `json:"delayMs"`
	AmountETH    float64  `json:"amountETH"`
}

// SellBundleRequest is the payload for POST /api/sell/bundle.
type SellBundleRequest struct {
	TokenAddress string   `json:"tokenAddress"`
	Wallets      []string `json:"wallets"`
	SellPercent  float64  `json:"sellPercent"`
	BribeETH     float64  `json:"bribeETH"`
}

// DevingArmRequest is the payload for POST /api/deving/arm.
type DevingArmRequest struct {
	Mode            string   `json:"mode"`
	PresetID        string   `json:"presetId"`
	TwitterKeywords []string `json:"twitterKeywords"`
	TrackedDevs     []string `json:"trackedDevs"`
}

// TokenCreateRequest is the payload for POST /api/token/create.
type TokenCreateRequest struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"imageUrl"`
	Supply       uint64  `json:"supply"`
	LiquidityETH float64 `json:"liquidityETH"`
}

// SnipeResponse is returned by the snipe/sell endpoints.
type SnipeResponse struct {
	Success  bool     `json:"success"`
	TxHashes []string `json:"txHashes,omitempty"`
	Error    string   `json:"error"`
}

// TokenCreateResponse is returned by POST /api/token/create.
type TokenCreateResponse struct {
	Success      bool   `json:"success"`
	TokenAddress string `json:"tokenAddress,omitempty"`
	PairAddress  string `json:"pairAddress,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
	Error        string `json:"error"`
}
