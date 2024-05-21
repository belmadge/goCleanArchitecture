package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "user"
    password = "password"
    dbname   = "orders_db"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer db.Close()

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            item VARCHAR(50),
            quantity INT,
            price DECIMAL
        )
    `)
    if err != nil {
        log.Fatalf("Unable to execute migration: %v\n", err)
    }

    log.Println("Migration completed successfully!")
}
