package users

import (
	"database/sql"
	"devbook-api/app/entity"
	"fmt"
)

type GetUserReposistory struct {
	db *sql.DB
}

func NewGetRepository(db *sql.DB) GetUserReposistory {
	return GetUserReposistory{db}
}

func (repository GetUserReposistory) GetAll(nameOrNick string) ([]entity.UsersEntity, error) {
	nameOrNick = fmt.Sprintf("%%%%%s", nameOrNick) // esses %% sao para escate(deixar escapar %nameOrNick%
	rows, erro := repository.db.Query(
		"SELECT id_user, user_name, nick, email, created_at, updated_at FROM users where user_name LIKE $1 OR nick LIKE $2",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []entity.UsersEntity
	for rows.Next() {
		var user entity.UsersEntity
		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.Created_at,
			&user.Updated_at,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}
