package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalidade).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r),
	))
}
