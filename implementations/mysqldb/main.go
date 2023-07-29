package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tenlisboa/go-built-in-examples/implementations/mysqldb/pkg/repository"
)

func getArgs() (string, string) {
  args := os.Args[1:]

  if len(args) != 2 {
    fmt.Println("Usage: go run main.go <username> <password>")
    os.Exit(1)
  }

  username := args[0]
  password := args[1]

  return username, password
}

func main() {
  username, password := getArgs()

  db, err := sql.Open("mysql", "myuser:mypass@tcp(localhost:3306)/mydb?parseTime=true")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  userRepo := repository.NewUserRepository(db) 
  userExists, _ := userRepo.GetUserByUsername(db, username)
  if userExists {
    fmt.Println("User already exists")
    return
  }

  userRepo.CreateUser(db, username, password)
  users := userRepo.GetAllUsers(db)

  fmt.Println(users)

  err = db.Ping()
}


