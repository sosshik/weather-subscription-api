package repository

import "time"

type Subscription struct {
	Email       string    `db:"email"`
	City        string    `db:"city"`
	Frequency   string    `db:"frequency"`
	Token       string    `db:"token"`
	IsConfirmed bool      `db:"is_confirmed"`
	CreatedAt   time.Time `db:"created_at"`
}
