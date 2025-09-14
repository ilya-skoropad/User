package repository

import (
	"database/sql"
	"ilya-skoropad/user/internal/entity"

	_ "github.com/lib/pq"
)

type UserRepository interface {
	CheckExistsByLoginOrMail(login string, email string) error
	Save(entity.User) error
}

type userRepository struct {
	connection *sql.DB
}

func (r userRepository) CheckExistsByLoginOrMail(login string, email string) error {
	var guid string
	err := r.connection.QueryRow("SELECT a.guid FROM app.user a WHERE login = $1 OR email = $2 LIMIT 1", login, email).Scan(&guid)
	if err != nil {
		return err
	}

	return err
}

func (r *userRepository) Save(user entity.User) error {
	statment := `insert into app."user" ("state", "role", "nickname", "login", "email", "password")
		values (
			(select id from app.state s where s."name" = 'Active'),
			(select id from app.role s where s."name" = 'User'),
			$1, $2, $3, $4)`

	_, err := r.connection.Exec(statment, user.Nickname, user.Login, user.Email, user.Password)
	return err
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}
