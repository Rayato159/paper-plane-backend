package repositories

import (
	"errors"

	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) users.Repository {
	return &usersRepository{db: db}
}

func (r *usersRepository) Register(m *models.RegisterUser) (string, error) {
	query := `INSERT INTO users (id, username, password) VALUES (?, ?, ?)`

	_, err := r.db.Exec(query, m.Id, m.Username, m.Password)
	if err != nil {
		return "", err
	}
	return m.Id, nil
}

func (r *usersRepository) GetUserResponse(id string) (*models.UserResponse, error) {
	query := `SELECT id, username FROM users WHERE id = ? LIMIT 1`

	var user models.UserResponse
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) DeleteUser(id string) error {
	query := `
DELETE FROM accounts WHERE user_id = ? LIMIT 1;
`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	query = `
DELETE FROM users WHERE id = ? LIMIT 1;
`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *usersRepository) GetUserPassword(id string) (*string, error) {
	query := `
SELECT
password
FROM users
WHERE id = ?
LIMIT 1
`

	password := new(string)
	if err := r.db.Get(password, query, id); err != nil {
		return nil, err
	}
	return password, nil
}

func (r *usersRepository) UserJWTPayloadCheck(payload *models.UserJWTPayload) (*models.User, error) {
	query := `
SELECT
u.id,
u.username
FROM users u
WHERE u.id = ?
AND u.username = ?
LIMIT 1
`

	user := new(models.User)
	if err := r.db.Get(user, query, payload.UserId, payload.Username); err != nil {
		return nil, errors.New("error, user not found.")
	}
	return user, nil
}
