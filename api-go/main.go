package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* ---------- Domain ---------- */

type Order struct {
	ID         int       `json:"id"`
	Customer   string    `json:"customer"`
	Product    string    `json:"product"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

var (
	store  = make(map[int]Order)
	nextID = 1
	mu     sync.Mutex
)

/* ---------- Handlers ---------- */

// POST /orders
func createOrder(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Customer  string  `json:"customer"`
		Product   string  `json:"product"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unit_price,omitempty"` // optional
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	order := Order{
		ID:         nextID,
		Customer:   input.Customer,
		Product:    input.Product,
		Quantity:   input.Quantity,
		TotalPrice: float64(input.Quantity) * input.UnitPrice,
		CreatedAt:  time.Now(),
	}
	store[nextID] = order
	nextID++
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GET /orders/{id}
func getOrder(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/orders/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	order, ok := store[id]
	mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

/* ---------- main ---------- */

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createOrder(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// route for /orders/{id}
	http.HandleFunc("/orders/", getOrder)

	fmt.Println("Order API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
