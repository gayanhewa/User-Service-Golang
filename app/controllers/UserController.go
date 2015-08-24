package controllers

import (
  "fmt"
  "./../repo"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "./../models"
  "gopkg.in/bluesuncorp/validator.v7"
  "./../custom"
)

var validate *validator.Validate
var response *custom.Response

type UserController struct {

}

type Auth struct {
  Username string
  Password string
}

// Get All Users
func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {

  var uRepo = repo.UserRepository{}
  res,err := uRepo.GetAll()

  // Error occured
  if err != nil {
    response.Format(w, r, true, 400, err)
  }

  response.Format(w, r, false, 200, res)
}

//Get One By ID
func (u *UserController) GetOne(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var uRepo = repo.UserRepository{}
  var user,err = uRepo.GetOne(vars)

  res, err := json.Marshal(user)

  if err != nil {
    w.Write([]byte("Error"))
  }

  w.Write([]byte(res))

}

// Destroy user
func (u *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var uRepo = repo.UserRepository{}
  var user,err = uRepo.Destroy(vars)

  res, err := json.Marshal(user)

  if err != nil {
    w.Write([]byte("Error"))
  }

  w.Write([]byte(res))

}

// Update User
func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var uRepo = repo.UserRepository{}
  var user,err = uRepo.Update(vars)

  res, err := json.Marshal(user)

  if err != nil {
    w.Write([]byte("Error"))
  }

  w.Write([]byte(res))

}

// Create User
func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Imported and not used")
    var user models.User
    var uRepo = repo.UserRepository{}

    decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&user)

    if err != nil {
     response.Format(w, r, true, 417, err)
     return
    }

    config := validator.Config{
                TagName:         "validate",
                ValidationFuncs: validator.BakedInValidators,
    }

    validate = validator.New(config)

    errs := validate.Struct(user);

    if errs != nil {
     response.Format(w, r, true, 417, errs)
     return
    }
 
    created, create_err := uRepo.Create(user)

    if create_err != nil {
      response.Format(w, r, true, 418, create_err)
      return
    }

    if created {
     response.Format(w, r, false, 200, user)
     return
    }

}

func (u *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {

    var authuser models.AuthUser
    var uRepo = repo.UserRepository{}

    decoder := json.NewDecoder(r.Body)

    decoder.Decode(&authuser)

    user,err := uRepo.Authenticate(authuser)

    if err != nil {
     response.Format(w, r, true, 404, false)
     return
    }
      
    response.Format(w, r, false, 200, user)
    return

}
