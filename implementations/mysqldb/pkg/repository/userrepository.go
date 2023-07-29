package repository

import (
	"database/sql"
	"time"
)

type User struct {
  Id        int
  Username  string
  Password  string
  CreatedAt time.Time
}

type UserRepository struct {
  db *sql.DB
}

func createTableUsers(db *sql.DB) {
  query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

  if _, err := db.Exec(query); err != nil {
    panic(err)
  }
}

func NewUserRepository(db *sql.DB) *UserRepository {
  userRepo := UserRepository{
    db: db,
  }

  createTableUsers(db)

  return &userRepo
}

func (ur *UserRepository) GetUserByUsername(db *sql.DB, username string) (bool, User) {
  query := `
    SELECT id, username, password, created_at
    FROM users
    WHERE username = ?`

  var user User

  err := db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
  if err != nil {
    return false, user
  }

  return true, user
}

func (ur *UserRepository) GetAllUsers(db *sql.DB) []User {
  rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
  if err != nil {
    panic(err)
  }

  defer rows.Close()

  var users []User

  for rows.Next() {
    var u User
    err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
    if err != nil {
      panic(err)
    }
    users = append(users, u)
  }

  return users
}

func (ur *UserRepository) CreateUser(db *sql.DB, username string, password string) {
  query := `
    INSERT INTO users (username, password, created_at)
    VALUES (?, ?, ?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(username, password, time.Now())
  if err != nil {
    panic(err)
  }
}
