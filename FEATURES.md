# Features — RobinHood EVM Dev Sniper

Full API reference with request/response examples for all endpoints.

> All endpoints require `Authorization: Bearer <your-license-key>` header.
> Base URL: `http://your-server:8080`

**Tier legend:** 🟢 **Core** (1 ETH) · 🔵 **Full Suite** (1.5 ETH, includes everything in Core)

| # | Feature | Tier |
|---|---------|------|
| 1 | Multi-Wallet Sniping | 🟢 Core |
| 2 | Flashbots Bundle Protection | 🟢 Core |
| 3 | Volume Bot | 🟢 Core |
| 4 | WebSocket Monitoring | 🟢 Core |
| 5 | Token Analysis | 🟢 Core |
| 6 | Bundle Sell | 🟢 Core |
| 7 | AI Deving Autopilot | 🔵 Full Suite |
| 8 | Token Creator | 🔵 Full Suite |

---

## 1. Multi-Wallet Sniping

Coordinate simultaneous buys across multiple wallets the moment fresh liquidity is added on Uniswap. Each wallet sends an independent RPC transaction in parallel, maximizing fill probability.

**Endpoint:** `POST /api/snipe/fire`

```bash
curl -X POST http://localhost:8080/api/snipe/fire \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenAddress": "0xTokenContract...",
    "wallets": ["0xwallet1", "0xwallet2", "0xwallet3"],
    "buyPercent": 25.0,
    "slippageBps": 500,
    "maxGasGwei": 40
  }'
```

**Response:**
```json
{
  "success": true,
  "txHashes": [
    "0x5KtP...tx1",
    "0x7mNq...tx2",
    "0x9pRz...tx3"
  ],
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenAddress` | string | ERC-20 token contract address |
| `wallets` | []string | Wallet addresses to snipe with |
| `buyPercent` | float64 | % of each wallet's ETH balance to spend |
| `slippageBps` | uint16 | Slippage tolerance in basis points (500 = 5%) |
| `maxGasGwei` | uint64 | Max gas price cap in gwei |

---

## 2. Flashbots Bundle Protection

Submit buy transactions as a Flashbots bundle for MEV protection. Bundles are included atomically — all wallets buy in the same block, preventing sandwich attacks and front-running via the public mempool.

**Endpoint:** `POST /api/snipe/bundle-fire`

```bash
curl -X POST http://localhost:8080/api/snipe/bundle-fire \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenAddress": "0xTokenContract...",
    "wallets": ["0xwallet1", "0xwallet2"],
    "bribeETH": 0.01,
    "rpcFallbackDelayMs": 500
  }'
```

**Response:**
```json
{
  "success": true,
  "bundleHash": "0xbundle_abc123...",
  "txHashes": ["0x5KtP...tx1", "0x7mNq...tx2"],
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenAddress` | string | ERC-20 token contract address |
| `wallets` | []string | Wallet addresses for bundle |
| `bribeETH` | float64 | Builder bribe in ETH (higher = faster inclusion) |
| `rpcFallbackDelayMs` | int64 | Fallback to RPC if bundle not included within N ms |

---

## 3. Volume Bot

Generate automated buy/sell cycles across wallets to create organic-looking trading volume on any ERC-20 pair. Useful for maintaining chart momentum.

**Endpoint:** `POST /api/volume/start`

```bash
curl -X POST http://localhost:8080/api/volume/start \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenAddress": "0xTokenContract...",
    "wallets": ["0xwallet1", "0xwallet2", "0xwallet3"],
    "cycles": 10,
    "delayMs": 2000,
    "amountETH": 0.1
  }'
```

**Response:**
```json
{
  "success": true,
  "jobId": "vol_job_xyz789",
  "estimatedDurationMs": 20000,
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenAddress` | string | ERC-20 token contract address |
| `wallets` | []string | Wallets to rotate through |
| `cycles` | int | Number of buy/sell cycles to execute |
| `delayMs` | int64 | Delay between cycles in milliseconds |
| `amountETH` | float64 | ETH amount per buy transaction |

---

## 4. WebSocket Monitoring

Subscribe to real-time transaction events for any ERC-20 pair. Receive live price updates, buy/sell events, take-profit triggers, and stop-loss alerts.

**Endpoint:** `GET /api/monitor/ws` (WebSocket upgrade)

```javascript
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/api/monitor/ws', {
  headers: { 'Authorization': 'Bearer YOUR_LICENSE_KEY' }
});

ws.on('open', () => {
  // Subscribe to a token
  ws.send(JSON.stringify({
    action: 'subscribe',
    tokenAddress: '0xTokenContract...'
  }));
});

ws.on('message', (data) => {
  const event = JSON.parse(data);
  // event.type: 'trade' | 'take_profit' | 'stop_loss' | 'price_update'
  console.log(`[${event.type}] ${event.tokenAddress} — ${event.action} ${event.amountETH} ETH`);
});
```

**Event payload:**
```json
{
  "type": "trade",
  "tokenAddress": "0xTokenContract...",
  "action": "buy",
  "amountETH": 1.5,
  "priceUSD": 0.00042,
  "marketCap": 42000,
  "timestamp": "2026-04-01T12:00:00Z",
  "txHash": "0x5KtP...txhash"
}
```

---

## 5. Token Analysis

Fetch current pair data for any ERC-20 token — liquidity, reserves, buy/sell tax, and deployer info.

**Endpoint:** `GET /api/token/:address`

```bash
curl http://localhost:8080/api/token/0xTokenContract... \
  -H "Authorization: Bearer YOUR_LICENSE_KEY"
```

**Response:**
```json
{
  "tokenAddress": "0xTokenContract...",
  "liquidityETH": 12.3,
  "marketCapETH": 42.5,
  "buyTaxBps": 300,
  "sellTaxBps": 300,
  "liquidityLocked": true,
  "deployer": "0xDeployerWallet..."
}
```

---

## 6. Bundle Sell

Execute coordinated sells across multiple wallets via Flashbots bundle. Ensures all wallets exit in the same block to minimize price impact.

**Endpoint:** `POST /api/sell/bundle`

```bash
curl -X POST http://localhost:8080/api/sell/bundle \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenAddress": "0xTokenContract...",
    "wallets": ["0xwallet1", "0xwallet2"],
    "sellPercent": 100.0,
    "bribeETH": 0.005
  }'
```

**Response:**
```json
{
  "success": true,
  "txHashes": ["0x5KtP...tx1", "0x7mNq...tx2"],
  "ethReceived": 3.42,
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenAddress` | string | ERC-20 token contract address |
| `wallets` | []string | Wallets holding tokens to sell |
| `sellPercent` | float64 | % of each wallet's token balance to sell |
| `bribeETH` | float64 | Builder bribe in ETH |

---

## 7. AI Deving Autopilot 🔵 Full Suite

Arm the autopilot to deploy from your saved presets automatically, triggered by real-time Twitter/X signals or a watched top dev deploying. In semi-manual mode it surfaces the relevant tweet for one-click deploy; in full autopilot it fires on its own.

**Endpoint:** `POST /api/deving/arm`

```bash
curl -X POST http://localhost:8080/api/deving/arm \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "mode": "autopilot",
    "presetId": "preset_fast_entry",
    "twitterKeywords": ["launch", "live now"],
    "trackedDevs": ["0xDevWallet1...", "0xDevWallet2..."]
  }'
```

**Response:**
```json
{
  "success": true,
  "armed": true,
  "mode": "autopilot",
  "watching": { "keywords": 2, "devs": 2 },
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `mode` | string | `semi` (one-click confirm) or `autopilot` (auto-deploy) |
| `presetId` | string | Saved snipe preset to fire with |
| `twitterKeywords` | []string | X/Twitter keywords that trigger a deploy |
| `trackedDevs` | []string | Dev wallets to mirror on deploy |

---

## 8. Token Creator 🔵 Full Suite

Deploy an ERC-20 token and open a Uniswap pool with full metadata in a single call.

**Endpoint:** `POST /api/token/create`

```bash
curl -X POST http://localhost:8080/api/token/create \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Token",
    "symbol": "MYTKN",
    "description": "Launched with RobinHood EVM Dev Sniper",
    "imageUrl": "https://.../logo.png",
    "supply": 1000000000,
    "liquidityETH": 0.5
  }'
```

**Response:**
```json
{
  "success": true,
  "tokenAddress": "0xNewTokenContract...",
  "pairAddress": "0xNewUniswapPair...",
  "txHash": "0x5KtP...txhash",
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Token name |
| `symbol` | string | Token ticker |
| `description` | string | Token description / metadata |
| `imageUrl` | string | Token image URL |
| `supply` | uint64 | Total token supply to mint |
| `liquidityETH` | float64 | Initial ETH paired into the Uniswap pool |

---

## Authentication

All endpoints require a valid license key:

```
Authorization: Bearer YOUR_LICENSE_KEY
```

License keys are delivered via email after purchase at [memesnipe.fun](https://memesnipe.fun).

---

## Error Responses

```json
{
  "success": false,
  "error": "insufficient ETH balance in wallet 0xwallet1"
}
```

Common error codes:
- `401 Unauthorized` — invalid or expired license key
- `400 Bad Request` — missing required fields or invalid parameters
- `429 Too Many Requests` — rate limit exceeded
- `500 Internal Server Error` — RPC or blockchain error (retryable)
