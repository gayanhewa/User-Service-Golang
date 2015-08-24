package models

import (
  "gopkg.in/mgo.v2/bson"
  "golang.org/x/crypto/bcrypt"
)

type AuthUser struct{
	Username string
	Password string
}

type User struct {
  Id  bson.ObjectId `_id,omitempty`
  FirstName string `validate:"required" json:"first_name"`
  LastName string `validate:"required" json:"last_name"`
  Email string `validate:"required,email" json:"email"`
  Username string `validate:"required" json:"username"`
  Password string `validate:"required" json:"-"`
}

func (u *User) SetPassword(password string) {
  hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
      panic(err) //this is a panic because bcrypt errors on invalid costs
  }
  u.Password = string(hpass)
}

func (u *User) CheckPassword(password string) bool{
  err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
  if err != nil {
      return false
  }

  return true
}