package models

import (

	// "github.com/dgrijalva/jwt-go"
	"time"
)

// customer type to add new customer it contains user table attributes
type Customer struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone int `json:"phone"`
	Password string `json:"password"`
	Place string `json:"place"`
	Dob int `json:"dob"`
	Occupation string `json:"occupation"`
	
}

//to fetch customer data from database 
type CustomerData struct{
	Userid int `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone int `json:"phone"`
	Place string `json:"place"`
	Dob time.Time `json:"dob"`
	Occupation string `json:"occupation"`
}

//this typw is to add new account to an existing user
type Account struct{
	Userid int `json:"user_id"`
	Acc_no int `json:"acc_no"`
	Acc_type string `json:"acc_type"`
	Balance float32 `json:"balance"`
}

//this is to add new transaction
type Transaction struct{
	Userid int `json:"userid"`
	Acc_id int `json:"acc_id"`
	Amount int `json:"amount"`
	Category string `json:"category"`
	Type string `json:"type"`
	Description string `json:"description"`
}

//to retuen transaction details
type TransactionData struct{
	Trans_id int `json:"transactionid"`
	Userid int `json:"userid"`
	Acc_id int `json:"acc_id"`
	Amount int `json:"amount"`
	Category string `json:"category"`
	Type string `json:"type"`
	Description string `json:"description"`
	Time time.Time `json:"transaction_time"`
}

//to store data from user to login validation
// type credentials struct{
// 	Email string `json:"email"`
// 	Password string `json:"password"`
// }

// //to store jwt token claims
// type claims struct{
// 	Userid int `json:"user_id"`
// 	jwt.StandardClaims
	
// }
