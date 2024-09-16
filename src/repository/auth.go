package repository

import "github.com/jmoiron/sqlx"

type AuthRepository struct {
	Db *sqlx.DB
}

func NewAuthRepository(Db *sqlx.DB) AuthRepository {
	return AuthRepository{Db: Db}
}

func (repo *AuthRepository) getUserPassword(email string) string {
	var password string

	query := "SELECT password FROM users WHERE email=$1"
	err := repo.Db.Get(&password, query, email)
	if err != nil {
		//handle error
	}
	return password
}

// func (repo *AuthRepository) GetUserRefreshToken(userId)
