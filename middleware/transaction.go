package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"transhabit/models"
	"github.com/gorilla/mux"
)


func AddTransaction(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	var trans models.Transaction

	json.NewDecoder(r.Body).Decode(&trans)

	statement:= "INSERT INTO transaction (userid,amount,category,type,time) values ($1,$2,$3,$4,$5)"

	res,err:=db.Exec(statement,trans.Userid,trans.Amount,trans.Category,trans.Type,time.Now())

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

func TransactionbyId(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	input:=mux.Vars(r)

	in,err:=strconv.Atoi(input["id"])

	if err!=nil{
		log.Fatalf("error while input convert -%v",err)
	}

	statement:="SELECT * FROM transaction WHERE id=$1"

	row:=db.QueryRow(statement,in)

	if err!=nil{
		log.Fatalf("error occured during querying- %v",err)
	}
	var trans models.Transaction

	errors:=row.Scan(&trans.Id,&trans.Userid,&trans.Amount,&trans.Category,&trans.Type,&trans.Time)

	if errors!=nil{
		log.Fatalf("error while converting data -%v",errors)
	}

	json.NewEncoder(w).Encode(trans)

}

func ListTransactions(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	statement:="SELECT * FROM transaction"

	rows,err:=db.Query(statement)

	if err !=nil{
		log.Fatalf("error occured while querying -%v",err)
	}

	var transactions []models.Transaction

	for rows.Next(){
		var trans models.Transaction

		err:=rows.Scan(&trans.Id,&trans.Userid,&trans.Amount,&trans.Category,&trans.Type,&trans.Time)

		if err!=nil{
			log.Fatalf("error while loading each row - %v",err)
		}

		transactions=append(transactions, trans)
	}

	json.NewEncoder(w).Encode(transactions)

}


func DeleteTransaction(w http.ResponseWriter,r *http.Request){
	db:=ConnectDB()
	defer db.Close()

	param:=mux.Vars(r)

	todel,_:=strconv.Atoi(param["id"])

	statement:="DELETE FROM transaction WHERE id=$1"

	res,err:=db.Exec(statement,todel)

	if err!=nil{
		log.Fatalf("error while executing query -%v",err)
	}

	mes,_:=res.RowsAffected()

	mes_to_send:=fmt.Sprintf("No. of rows deleted: %v",mes)

	out:=response{Message: mes_to_send}

	json.NewEncoder(w).Encode(out)
}
