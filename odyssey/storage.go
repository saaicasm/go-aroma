package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	godotenv.Load()
	connStr := fmt.Sprintf("user=postgres dbname=postgres password=%s sslmode=disable", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {

	return s.CreateAccountTable()

}

func (s *PostgresStore) CreateAccountTable() error {

	query := `
		CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number SERIAL UNIQUE,
		balance DECIMAL(10,2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAxMP
	);`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	query := `
		insert into account 
		(first_name, last_name, number, balance, created_at)
		values($1, $2, $3, $4, $5)
	`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.Lastname,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", resp)

	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		fmt.Println("Error is in select statement")
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil

}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.Lastname,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)

	if err != nil {
		return err
	}

	fmt.Printf("Deleted %d", id)

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	rows, err := s.db.Query("select * from account where id = $1", id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found ", id)
}
