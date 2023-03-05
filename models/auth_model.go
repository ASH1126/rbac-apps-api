package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"rbac-api/config/db"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
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

	query := `SELECT id, email,	first_name,	last_name, password, user_active, created_at, updated_at FROM users WHERE email = $1`
	rows, err := dbEngine.QueryString(query, email)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("invalid username/password")
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

func (u *User) PasswordMatches(planText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(planText))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token))
	token.TokenHash = hash[:]

	return token, nil
}

func Insert(token Token, u User) error {
	dbEngine := db.ConnectDB()

	// delete any existing tokens
	stmt := "delete from tokens where user_id =?"
	_, err := dbEngine.Exec(stmt, token.UserID)
	if err != nil {
		return err
	}

	// we assign the email value, just to be safe, in case it was
	// not done in the handler that calls this function
	token.Email = u.Email

	// insert the new token
	stmt = `insert into tokens (user_id, email, token, token_hash, created, updated, expiry)
		values ($1, $2, $3, $4, $5, $6, $7)`
	_, err = dbEngine.Exec(stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		time.Now(),
		time.Now(),
		token.Expiry,
	)
	if err != nil {
		return err
	}

	return nil
}
