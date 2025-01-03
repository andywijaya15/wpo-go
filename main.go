package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"wpo-go/config"
)

type ChangedPurchaseOrder struct {
	COrderID uint
	Created  time.Time
}

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	response, err := http.Get("http://localhost:8080/v1/get-changed-purchase-orders-concurrency")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch data. HTTP Status: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var changedPurchaseOrders []ChangedPurchaseOrder
	err = json.Unmarshal(body, &changedPurchaseOrders)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	for _, order := range changedPurchaseOrders {
		fmt.Printf("Order ID: %d, Created: %s\n", order.COrderID, order.Created.Format(time.DateOnly))
	}
}
