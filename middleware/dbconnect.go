package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB{
	// dsn := "postgresql://kaviraj:hscSNtKlb6M8BDYlmVSbxQ@kaviraj-test-4006.j77.aws-ap-south-1.cockroachlabs.cloud:26257/transhabit?sslmode=verify-full"
	db_pass:=os.Getenv("PG_PASSWORD")
	dsn :=fmt.Sprintf("postgres://postgres:%v@localhost/transhabit?sslmode=disable",db_pass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("Database connected..")

	return db

}