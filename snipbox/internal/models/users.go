package models

import (
	"database/sql"
	"time"
)

type users struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name string, password string, email string) error {
	return nil
}

func (m *UserModel) Authenticate(email string, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(email string) (bool, error) {
	return false, nil
}
