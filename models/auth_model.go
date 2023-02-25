package models

import (
	"fmt"
	"rbac-api/config/db"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Token     Token     `json:"token"`
}

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Expiry    time.Time `json:"expiry"`
}

func GetByEmail(email string) (*User, error) {
	dbEngine := db.ConnectDB()

	// query := `SELECT id, email,	first_name,	last_name, password, user_active, created_at, updated_at FROM users WHERE email = $1`
	// var user User

	// row, err := dbEngine.QueryInterface(query, email)
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Println(dbEngine)

	// err = row.Scan(
	// 	&user.ID,
	// 	&user.Email,
	// 	&user.FirstName,
	// 	&user.LastName,
	// 	&user.Password,
	// 	&user.Active,
	// 	&user.CreatedAt,
	// 	&user.UpdateAt,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
