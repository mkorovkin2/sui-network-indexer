package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// API endpoint to get wallet details
	r.GET("/getWalletDetails", func(c *gin.Context) {
		address := c.Query("address")
		if address == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "wallet address is required"})
			return
		}

		// Fetch wallet details by querying the local Sui node
		details, err := GetWalletDetails(address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, details)
	})

	// Start the server on port 8080
	r.Run(":8080")
}
