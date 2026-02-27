package model

import "time"

type User struct {
    ID           string    `db:"id"`
    Email        string    `db:"email"`
    Username     string    `db:"username"`
    PasswordHash string    `db:"password_hash"`
    CreatedAt    time.Time `db:"created_at"`
}

