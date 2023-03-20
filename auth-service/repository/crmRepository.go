package repository

import (
	"log"
	"os"

	"github.com/a-rajgaria/HostelProject/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func DBConnect() {
	db_url := os.Getenv("DB_URL")

	log.Println( db_url)
	
	db_connection, err := sqlx.Open("postgres", db_url)
	
	if err != nil {
		log.Fatalln("error", err)
		log.Fatalln("Not able to connect with db!")
	}
	
	DB = db_connection

	if err:=db_connection.Ping(); err!=nil{
		log.Println("Error", err)
		panic("DB is not connected !!")
	}

	log.Println("DB is connected!")

}

func DBCreate(customer models.Customer) models.Customer{
	query := "INSERT INTO customer (name, email, password, mobile) VALUES ($1, $2, $3, $4) RETURNING id, name, email, mobile ;"
	row, err1 := DB.Queryx(query,
		customer.Name,
		customer.Email,
		customer.Password,
		customer.Mobile,
	)
	if err1 != nil {
		log.Println("Error", err1)
		log.Println("Not able to Insert customer into Database")
	}
	newCustomer := models.Customer{}
	for row.Next() {
		err2 := row.StructScan(&newCustomer)
		if err2 != nil {
			log.Println("Error", err2)
			log.Println("Not able to create new customer")
		}
	}
	return newCustomer
}

func DBGetByEmail(email string) models.Customer{
	
	customer := models.Customer{}
	query := "SELECT id, email, password FROM customer WHERE email = $1"
	err1 := DB.Get(&customer, query, email)
	if err1 != nil {
		log.Println("Error", err1)
	}
	return customer
}