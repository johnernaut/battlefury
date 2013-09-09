package db

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

type connection struct {
    Db *sql.DB
}

type Model interface {
    Find() string
}

func Find(m Model) string {
    return m.Find()
}

var Connection = new(connection)
var Models = []Model{}

func Register(m Model) {
    Models = append(Models, m)
}

func Connect() {
    db, err := sql.Open("mysql", "root:@/go_development")

    if err !=  nil {
        log.Panic(err)
    }

    Connection.Db = db

    err = Connection.Db.Ping()

    if err != nil {
        log.Panic(err)
    }
}
