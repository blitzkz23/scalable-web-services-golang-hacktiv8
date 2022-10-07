package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "scalable-web-services-golang-hacktiv8/module_8/swaggo/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// ! Untuk membuka docs http://localhost:8080/swagger/index.html

// Order represents the model for an order
type Order struct {
	OrderID     string    `json:"orderId" example:"1"`
	CusomerName string    `json:"customerName" example:"John Doe"`
	OrderedAt   time.Time `json:"orderedAt" example:"2020-01-01T21:21:46+00:00"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ItemID      string `json:"itemId" example:"1"`
	Description string `json:"description example A book"`
	Quantity    int    `json:"quantity" example:"1"`
}

// @title ORDERS API
// @version 1.0
// @description This is a sample service for managing orders
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// Goriila Mux Router
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", createOrder).Methods("POST")
	// Get
	router.HandleFunc("/orders", getOrders).Methods("GET")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetOrders godoc
// @Summary Get orders
// @Description get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(orders)
}

// CreateOrder godoc
// @Summary Create order
// @Description create a new order with the input payload
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	// prevOrderId++
	// order.OrderID = strconv.Itoa(prevOrderId)
	// orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}
