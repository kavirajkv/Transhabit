package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"transhabit/models"

	"github.com/gorilla/mux"
)

type response struct{
	Message string `json:"message"`
}


func CreateCustomer(w http.ResponseWriter, r *http.Request ){
	db:=ConnectDB()
	defer db.Close()

	var cus models.Customer

	json.NewDecoder(r.Body).Decode(&cus)

	statement:= "INSERT INTO customer (name,email,phone,password,place,dob,occupation) values ($1,$2,$3,$4,$5,$6,$7)"

	res,err:=db.Exec(statement,cus.Name,cus.Email,cus.Phone,cus.Place,cus.Dob,cus.Occupation)

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

func CustomerbyId(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	input:=mux.Vars(r)

	in,err:=strconv.Atoi(input["id"])

	if err!=nil{
		log.Fatalf("error while input convert -%v",err)
	}

	statement:="SELECT user_id,name,email,phone,place,dob,occupation FROM customer WHERE user_id=$1"

	row:=db.QueryRow(statement,in)

	if err!=nil{
		log.Fatalf("error occured during querying- %v",err)
	}
	var cus models.CustomerData

	errors:=row.Scan(&cus.Userid,&cus.Name,&cus.Email,&cus.Phone,&cus.Place,&cus.Dob,&cus.Occupation)

	if errors!=nil{
		log.Fatalf("error while converting data -%v",errors)
	}

	json.NewEncoder(w).Encode(cus)

}


func ListCustomers(w http.ResponseWriter, r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	statement:="SELECT user_id,name,email,phone,place,dob,occupation FROM customer"

	rows,err:=db.Query(statement)

	if err !=nil{
		log.Fatalf("error occured while querying -%v",err)
	}

	var customers []models.CustomerData

	for rows.Next(){
		var cus models.CustomerData

		err:=rows.Scan(&cus.Userid,&cus.Name,&cus.Email,&cus.Phone,&cus.Place,&cus.Dob,&cus.Occupation)

		if err!=nil{
			log.Fatalf("error while loading each row - %v",err)
		}

		customers=append(customers, cus)
	}

	json.NewEncoder(w).Encode(customers)

}

func DeleteCustomer(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	param:=mux.Vars(r)

	todel,_:=strconv.Atoi(param["id"])

	statement:="DELETE FROM customer WHERE user_id=$1"

	res,err:=db.Exec(statement,todel)

	if err!=nil{
		log.Fatalf("error while executing query -%v",err)
	}

	mes,_:=res.RowsAffected()

	mes_to_send:=fmt.Sprintf("No. of rows deleted: %v",mes)

	out:=response{Message: mes_to_send}

	json.NewEncoder(w).Encode(out)
}

func UpdateCustomer(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	var cus models.CustomerData

	json.NewDecoder(r.Body).Decode(&cus)

	statement:="UPDATE customer SET name=$2  WHERE user_id=$1"

	row,err:=db.Exec(statement,cus.Userid,cus.Name)

	if err!=nil{
		log.Fatalf("error while executing query -%v",err)
	}

	res,_:=row.RowsAffected()

	msg:=fmt.Sprintf("No. rows affected: %v",res)

	mesg:=response{Message: msg}

	json.NewEncoder(w).Encode(mesg)

}
