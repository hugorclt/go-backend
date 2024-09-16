package repository

import (
	"database/sql"
	"go-backend/src/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

func (repo *UserRepository) GetUserById(userId string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id=$1"
	err := repo.DB.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
		}
	}
	return &user, nil
}
