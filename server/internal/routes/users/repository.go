package users

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := r.db.Exec(context.Background(), query, user.Email, user.Password)
	return err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
