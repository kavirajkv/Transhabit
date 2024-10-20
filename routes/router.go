package routes

import (
	"transhabit/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router:=mux.NewRouter()

	//routes for customer handling
	router.HandleFunc("/api/createcustomer",middleware.CreateCustomer).Methods("POST","OPTIONS")
	// router.HandleFunc("/api/updatecustomer",middleware.UpdateCustomer).Methods("PUT","OPTIONS")
	// router.HandleFunc("/api/deletecustomer",middleware.DeleteCustomer).Methods("DELETE","OPTIONS")
	// router.HandleFunc("/api/listcustomers",middleware.ListCustomers).Methods("GET","OPTIONS")
	router.HandleFunc("/api/customerbyId/{id}",middleware.CustomerbyId).Methods("GET","OPTIONS")

	// //routes for transaction
	// router.HandleFunc("/api/addtransaction",middleware.AddTransaction).Methods("POST","OPTIONS")
	// router.HandleFunc("/api/deletetransaction",middleware.DeleteTransaction).Methods("DELETE","OPTIONS")
	// router.HandleFunc("/api/listtransactions",middleware.ListTransactions).Methods("GET","OPTIONS")
	// router.HandleFunc("/api/transactionbyId",middleware.TransactionbyId).Methods("GET","OPTIONS")
	


	return router
}