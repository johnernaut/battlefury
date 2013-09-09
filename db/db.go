package db

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

type connection struct {
    Db *sql.DB
}

var Connection = new(connection)

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
