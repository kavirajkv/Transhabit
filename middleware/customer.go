package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transhabit/models"
)

type response struct{
	Message string `json:"message"`
}


func CreateCustomer(w http.ResponseWriter, r *http.Request ){
	db:=ConnectDB()
	defer db.Close()

	var cus models.Customer

	json.NewDecoder(r.Body).Decode(&cus)

	statement:= "INSERT INTO customer (name,place, age, occupation,balance) values ($1,$2,$3,$4,$5)"

	res,err:=db.Exec(statement,cus.Name,cus.Place,cus.Age,cus.Occupation,cus.Balance)

	if err!=nil{
		log.Fatalf("error executing query- %v",err)
	}

	mes,err:=res.RowsAffected()

	if err != nil{
		log.Fatalf("error while inserting - %v",err)
	}

	msg:=fmt.Sprintf("Insertion successful for %v",mes)

	res_msg:=response{Message: msg}

	json.NewEncoder(w).Encode(res_msg)

}

// func UpdateCustomer(){

// }

// func DeleteCustomer(){

// }

// func ListCustomers(){

// }

// func CustomerbyId(){

// }

