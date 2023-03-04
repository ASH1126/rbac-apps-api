package models

import (
	"rbac-api/config/db"
	"strconv"
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
	UpdateAt  time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}

func GetByEmail(email string) (*User, error) {
	dbEngine := db.ConnectDB()
	// loc, _ := time.LoadLocation("Asia/Jakarta")

	query := `SELECT id, email,	first_name,	last_name, password, user_active, created_at, updated_at FROM users WHERE email = $1`
	rows, err := dbEngine.QueryString(query, email)
	if err != nil {
		return nil, err
	}

	var user User

	for _, row := range rows {
		user.ID, _ = strconv.Atoi(row["id"])
		user.Email = row["email"]
		user.FirstName = row["first_name"]
		user.LastName = row["last_name"]
		user.Password = row["password"]
		user.Active, _ = strconv.Atoi(row["user_active"])
		user.CreatedAt, _ = time.Parse(time.RFC3339, row["created_at"])
		user.UpdateAt, _ = time.Parse(time.RFC3339, row["updated_at"])
	}

	return &user, nil
}
