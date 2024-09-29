package repository

import (
	"beauty_salon/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(input *entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (first_name, second_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"
	row := repo.db.QueryRow(query, input.FirstName, input.SecondName, input.Email, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *UserRepository) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := "SELECT * FROM users WHERE email=$1 AND password_hash=$2"
	err := repo.db.Get(&user, query, email, password)
	return user, err
}
