package repo

import (
  "fmt"
  "./../models"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "errors"
)

const (
  db_name = "myapp"
  url = "mongodb://dbuser:dbpass@localhost:57862/myapp"
)

type UserInterface interface {
  GetAll() ([]models.User, error)
  GetOne() (models.User, error)
  Destroy(id int) (bool, error)
  Update([]string) (bool, error)
  Create(user models.User) (bool, error)
  Authenticate(user models.User) bool
}

type UserRepository struct{}

func (u *UserRepository) GetAll() ([]models.User, error) {

  // Open connection to my mongodb instance
  session, err := mgo.Dial(url)

  if err != nil {
    return []models.User{},err
  }

  defer session.Close()

  result := []models.User{}
  users := session.DB(db_name).C("users")

  err = users.Find(nil).All(&result)

  return result,nil
}

// Get One
func (u *UserRepository) GetOne(vars map[string]string) (models.User, error) {

  // Open connection to my mongodb instance
  session, err := mgo.Dial(url)

  if err != nil {
    return models.User{},err
  }

  defer session.Close()

  result := models.User{}
  users := session.DB(db_name).C("users")

  err = users.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&result)

  return result,nil
}

// Destroy
func (u *UserRepository) Destroy(vars map[string]string) (bool, error) {

  // Open connection to my mongodb instance
  session, err := mgo.Dial(url)

  if err != nil {
    return false,err
  }

  defer session.Close()

  users := session.DB(db_name).C("users")

  err = users.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})

  if err != nil {
    return false, err
  }

  return true,nil
}

// Update
func (u *UserRepository) Update(vars map[string]string) (bool, error) {

  return false,nil
}

func (u *UserRepository) Create(user models.User) (bool, error) {

  // Open connection to my mongodb instance
  session, err := mgo.Dial(url)

  if err != nil {
    return false,err
  }

  defer session.Close()

  users := session.DB(db_name).C("users")
  user.SetPassword(user.Password)
  fmt.Println(user)
  err = users.Insert(user)

  if err != nil {
   return false,err
  }

  return true,nil
}

func (u *UserRepository) Authenticate(user models.AuthUser) (models.User,error) {
  
  // Open connection to my mongodb instance
  session, err := mgo.Dial(url)

  if err != nil {
    return models.User{}, err
  }

  defer session.Close()

  result := models.User{}
  users := session.DB(db_name).C("users")

  err = users.Find(bson.M{"username": user.Username}).One(&result)

  if result.CheckPassword(user.Password) {
    result.Password = ""
    return result,nil
  }

  return models.User{}, errors.New("Authentication Failed.")

}
