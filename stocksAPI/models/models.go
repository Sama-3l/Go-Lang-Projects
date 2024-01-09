package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}

func CreateConnection() *sql.DB {
	str, _ := os.Getwd()
	str = str + "\\middleware\\.env"
	err := godotenv.Load(str)

	if err != nil {
		log.Printf("Error loading .env file %v\n", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres")
	return db
}

func InsertStock(stock Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v \n", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func GetStock(id int64) (Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock Stock

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	err := db.QueryRow(sqlStatement, id).Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}

	return stock, err
}

func GetAllStocks() ([]Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stocks []Stock

	sqlStatement := `SELECT * FROM stocks `
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func UpdateStock(id int64, stock Stock) int64 {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stocksid=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected row %v", err)
	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}

func DeletedStock(id int64) int64 {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected row %v", err)
	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}
