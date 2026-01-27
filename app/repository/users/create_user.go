package users

import (
	"database/sql"
	"devbook-api/app/entity"
	"fmt"
)

type Usersrepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Usersrepository {
	return &Usersrepository{db}
}

func (repo Usersrepository) CreateUser(user entity.UsersEntity) (uint64, error) {
	var id_user uint64
	statement, err := repo.db.Prepare("insert into users(user_name, nick, email, password) values ($1, $2, $3, $4) RETURNING id_user")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = statement.QueryRow(user.Nome, user.Nick, user.Email, user.Password).Scan(&id_user)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	return uint64(id_user), nil
}
