package main

import (
  "fmt"
  "./app/controllers"
  "./app/routes"
)

var (
   UserController *controllers.UserController
   Dispatcher *routes.Dispatcher
 )
 
func main() {
  version := "1.0";
  fmt.Println("Starting version "+version)

  // Initialize the dispatcher
  Dispatcher.Init()
}

