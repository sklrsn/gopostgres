package store

import (
	"database/sql"
	"fmt"
)

//SayHello - Just for Checking
func SayHello() string {
	return "Hello sir!!"
}

//Sum - Add t
func Sum(a, b int) int {
	return a + b
}

//Product - Holds Product Attributes
type Product struct {
	ID    int
	Name  string
	Price float64
}

//Connection - store db Connection
type Connection struct {
	DB *sql.DB
}

//Persist - Store Product attributes in Database
func Persist(product *Product, connection *Connection) int {
	var err error
	sqlStatement := `INSERT INTO products(name,price) VALUES($1,$2) RETURNING id`
	id := 0
	err = connection.DB.QueryRow(sqlStatement, product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		panic("Insert data failed")
	}
	return id
}

//GetConnection - Get a Database Connection
func (connection *Connection) GetConnection() *sql.DB {
	return connection.DB
}

//Initialize - Create Database Connection
func (connection *Connection) Initialize(host, port, username, password, dbname string) (*sql.DB, error) {
	var err error
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s", host, port, username, password, dbname)
	connection.DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = connection.DB.Ping()
	if err != nil {
		panic(err)
	}
	return connection.DB, err
}
