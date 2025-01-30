package router

import (
	"github.com/gorilla/mux"
	"github.com/mebackend/profileapi/controller"

)

func Router() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/data", controller.CreateData).Methods("POST")

	router.HandleFunc("/api/users", controller.GetUserData).Methods("GET")

	return router
}

