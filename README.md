# 🧠 Sui Wallet Explorer

This is a small Go-based API service that connects to a **local Sui node** and returns a full set of on-chain data related to a specific wallet address.

---

## 🚀 Features

- Connects to a local Sui node via JSON-RPC.
- Queries all transactions initiated from a given wallet.
- Retrieves full transaction details including:
  - Inputs
  - Effects
  - Events
  - Object and Balance changes
- Returns results in human-readable JSON format.

---

## 🛠 Requirements

- Go 1.20+
- A local [Sui Full Node](https://docs.sui.io/build/full-node) running on `http://127.0.0.1:9000`
- Internet connection (to pull Gin)

---

## 📦 Installation

1. Clone the repo:

```bash
git clone https://github.com/yourusername/sui-wallet-service.git
cd sui-network-indexer
```

2. Download dependencies:

```bash
go mod tidy
```

3. Run the service:

```bash
go run main.go sui.go
```

---

## 🧪 Usage & Testing

Once the service is running, access the endpoint at:

```bash
GET http://localhost:8080/getWalletDetails?address=<WALLET_ADDRESS>
```

Example:

```bash
curl "http://localhost:8080/getWalletDetails?address=0x3a69d..."
```

This returns:

```bash
{
  "wallet": "0x3a69d...",
  "transactions": [
    {
      "digest": "...",
      "transaction": { ... },
      "events": [ ... ],
      "objectChanges": [ ... ],
      ...
    }
  ]
}
```

---

## 📚 Tips

* The service queries up to 100 recent transactions by default.
* You can extend it to handle pagination using the `cursor` in the Sui RPC response.
* Use a frontend to visualize the JSON if desired (Postman, Insomnia, or a browser extension).