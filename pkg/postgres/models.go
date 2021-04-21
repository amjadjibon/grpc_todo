package postgres

import (
	"time"
)

type Customer struct {
	tableName   struct{}  `pg:"customers"`
	Id          string    `pg:"id"`
	FirstName   string    `pg:"first_name"`
	LastName    string    `pg:"last_name"`
	PhoneNumber string    `pg:"phone_number"`
	Email       string    `pg:"email"`
	Password    string    `pg:"password"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
}