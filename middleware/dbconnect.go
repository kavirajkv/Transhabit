package middleware

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB{
//Password remove to make this repo public
	dsn := "postgresql://kaviraj:<kavidbpasswod>@kaviraj-test-4006.j77.aws-ap-south-1.cockroachlabs.cloud:26257/transhabit?sslmode=verify-full"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("Database connected..")

	return db

}