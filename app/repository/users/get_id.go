package users

import (
	"database/sql"
	"devbook-api/app/entity"
)

type GetIDUserReposistory struct {
	db *sql.DB
}

func NewGetIDUser(db *sql.DB) GetIDUserReposistory {
	return GetIDUserReposistory{
		db: db, // aqui poderia fazer igual ta nos outros, mas eu quis testar saber se vai da certo
	}
}

func (repository GetIDUserReposistory) GetID(ID uint64) (entity.UsersEntity, error) {
	rows, erro := repository.db.Query(
		"SELECT id_user, user_name, nick, email, created_at, updated_at FROM users WHERE id_user = $1",
		ID,
	)
	if erro != nil {
		return entity.UsersEntity{}, erro // vai mandar um usuario vazio
	}
	defer rows.Close()

	var user entity.UsersEntity
	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.Created_at,
			&user.Updated_at,
		); erro != nil {
			return entity.UsersEntity{}, erro
		}
	}

	return user, nil
}
