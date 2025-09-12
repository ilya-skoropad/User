package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository interface {
	FindGuidByLoginOrMail(login string, email string) (string, error)
}

type userRepository struct {
	connection *sql.DB
}

func (r userRepository) FindGuidByLoginOrMail(login string, email string) (string, error) {
	row, err := r.connection.Query("SELECT a.guid FROM app.user a WHERE login = $1 OR email = $2 LIMIT 1", login, email)
	if err != nil {
		return "", err
	}

	defer row.Close()

	var user string
	for row.Next() {
		err = row.Scan(&user)

		if err != nil {
			return "", err
		}
	}

	return user, err
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}
