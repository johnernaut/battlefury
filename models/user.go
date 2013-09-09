package models

import (
    "log"
	"github.com/johnernaut/battlefury/db"
)

type User struct {
    Username string `json:"username"`
    Email string `json:"email"`
}

func Find() []User {
    var username string
    var email string
    var users []User

    rows, err := db.Connection.Db.Query("select username,email from users")

    if err != nil {
        log.Panic(err)
    }
    
    for rows.Next() {
        rows.Scan(&username, &email)
        user := User{ Username: username, Email: email }
        users = append(users, user)
    }

    if err != nil {
        log.Panic(err)
    }

    return users
}
