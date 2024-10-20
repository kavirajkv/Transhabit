package main

import (
	"fmt"
	"net/http"
	"transhabit/routes"
)


func main(){
	r:=routes.Router()

	fmt.Println("Server running at port 8000")

	http.ListenAndServe(":8000",r)


}