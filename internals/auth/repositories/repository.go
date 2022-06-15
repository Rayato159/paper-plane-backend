package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/paper-plane/internals/auth"
	"github.com/paper-plane/internals/models"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepository{db: db}
}

func (r *authRepository) FindOneUser(username string) (*models.User, *models.UserCredentials, error) {
	query := `
SELECT 
u.id, 
u.username, 
u.password,
a.id AS account_id
FROM users u
LEFT JOIN accounts a
ON a.user_id = u.id
WHERE u.username = ? 
LIMIT 1
`

	var user models.User
	if err := r.db.Get(&user, query, username); err != nil {
		return nil, nil, err
	}

	credentials := new(models.UserCredentials)
	credentials.Id = user.Id.String()
	credentials.Username = user.Username
	credentials.AccountId = *user.AccountId

	return &user, credentials, nil
}

func (r *authRepository) CreateJwtToken(userId string, refreshToken string) error {
	query := `
UPDATE users
SET
refresh_token = ?
WHERE id = ?
LIMIT 1
`

	_, err := r.db.Exec(query, refreshToken, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) RefreshToken(refreshToken string) (*models.UserCredentials, error) {
	query := `
SELECT 
u.id, 
u.username, 
a.id AS account_id,
u.refresh_token
FROM users u
LEFT JOIN accounts a
ON a.user_id = u.id
WHERE u.refresh_token = ? 
LIMIT 1
`

	var credentials models.UserCredentials
	if err := r.db.Get(&credentials, query, refreshToken); err != nil {
		return nil, err
	}
	return &credentials, nil
}
