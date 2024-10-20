package models

import "time"


type Customer struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Place string `json:"place"`
	Age int `json:"age"`
	Occupation string `json:"occupation"`
	Balance int `json:"balance"`
}


type Transaction struct{
	Id int `json:"id"`
	Userid int `json:"userid"`
	Amount int `json:"amount"`
	Category string `json:"category"`
	Type string `json:"type"`
	Time time.Time `json:"time"`
}