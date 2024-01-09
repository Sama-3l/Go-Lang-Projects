package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stocksAPI/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateStock(w http.ResponseWriter, req *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(req.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode %v\n", err)
	}

	insertID := models.InsertStock(stock)

	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert to int %v", err)
	}

	stock, err := models.GetStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, req *http.Request) {
	stocks, err := models.GetAllStocks()

	if err != nil {
		log.Fatalf("Unable to get stocks. %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func DeleteStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Fatalf("Unable to convert to int. %v", err)
	}

	deletedStock := models.DeletedStock(id)
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/records %v", deletedStock)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert to int %v\n", err)
	}

	var stock models.Stock
	err = json.NewDecoder(req.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode request %v\n", err)
	}

	updateRows := models.UpdateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock update successfully. Total rows/records affected %v", updateRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
