package routes

import (
  "github.com/gorilla/mux"
  "net/http"
  "./../controllers"
  "fmt"
)

type Dispatcher struct {}

var (
  UserController *controllers.UserController
)

func (r *Dispatcher ) Init() {
  fmt.Println("Initialize the router")
  router := mux.NewRouter()

  router.StrictSlash(true)
  router.HandleFunc("/", profile).Methods("GET")
  // User Resource
  userRoutes := router.PathPrefix("/users").Subrouter()
  userRoutes.HandleFunc("/", UserController.GetAll).Methods("GET")
  userRoutes.HandleFunc("/", UserController.Create).Methods("POST")
  userRoutes.HandleFunc("/{id}", UserController.GetOne).Methods("GET")
  userRoutes.HandleFunc("/{id}", UserController.Destroy).Methods("DElETE")
  userRoutes.HandleFunc("/{id}", UserController.Update).Methods("PUT","PATCH")

  //Authenticate
  userRoutes.HandleFunc("/authenticate", UserController.Authenticate).Methods("POST")

  // bind the routes
  http.Handle("/", router)

  fmt.Println("Add the listner")

  //serve
  http.ListenAndServe(":9091", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("test"))
}
