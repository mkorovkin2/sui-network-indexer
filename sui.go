package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Local Sui node endpoint (assumes it's running on localhost:9000)
const SUI_NODE_URL = "http://127.0.0.1:9000"

// GetWalletDetails fetches all transaction details for a given wallet
func GetWalletDetails(address string) (map[string]interface{}, error) {
	// Step 1: Get transaction digests for the given wallet address
	txIDs, err := getTransactionsForAddress(address)
	if err != nil {
		return nil, err
	}

	// Step 2: Fetch full details for each transaction digest
	var txDetails []interface{}
	for _, tx := range txIDs {
		txDetail, err := getTransactionBlock(tx)
		if err == nil {
			txDetails = append(txDetails, txDetail)
		}
	}

	// Step 3: Return the full wallet object as JSON
	result := map[string]interface{}{
		"wallet":      address,
		"transactions": txDetails,
	}

	return result, nil
}

// getTransactionsForAddress queries the Sui node for transaction digests related to the address
func getTransactionsForAddress(address string) ([]string, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "sui_queryTransactionBlocks",
		"params": []interface{}{
			map[string]interface{}{
				"filter": map[string]interface{}{
					"FromAddress": address,
				},
				"options": map[string]interface{}{
					"showInput":          true,
					"showEffects":        true,
					"showEvents":         true,
					"showObjectChanges":  true,
					"showBalanceChanges": true,
				},
			},
			nil, // cursor
			100, // limit
			true, // descending order
		},
	}

	// Send request to Sui node
	body, err := sendRequest(payload)
	if err != nil {
		return nil, err
	}

	// Parse response and extract transaction digests
	var response struct {
		Result struct {
			Data []struct {
				Digest string `json:"digest"`
			} `json:"data"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	var txIDs []string
	for _, tx := range response.Result.Data {
		txIDs = append(txIDs, tx.Digest)
	}

	return txIDs, nil
}

// getTransactionBlock fetches full transaction block details for a given digest
func getTransactionBlock(digest string) (interface{}, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "sui_getTransactionBlock",
		"params": []interface{}{
			digest,
			map[string]interface{}{
				"showInput":          true,
				"showEffects":        true,
				"showEvents":         true,
				"showObjectChanges":  true,
				"showBalanceChanges": true,
			},
		},
	}

	body, err := sendRequest(payload)
	if err != nil {
		return nil, err
	}

	var response struct {
		Result interface{} `json:"result"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// sendRequest sends a generic JSON-RPC request to the Sui node
func sendRequest(payload map[string]interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(SUI_NODE_URL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
