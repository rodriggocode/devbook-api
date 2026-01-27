package users

import (
	"api-web/app/entity"
	"database/sql"
)

type UpatedUserRepository struct {
	db *sql.DB
}

func NewUpatedRepository(db *sql.DB) UpatedUserRepository {
	return UpatedUserRepository{db}
}

func (repository UpatedUserRepository) UpatedUser(ID uint64, user entity.UsersEntity) error {
	statemen, erro := repository.db.Prepare("UPDATE users SET user_name = $1, nick = $2, email = $3 WHERE id_user = $4")
	if erro != nil {
		return erro
	}
	defer statemen.Close()

	if _, erro = statemen.Exec(
		user.Nome,
		user.Nick,
		user.Email,
		ID,
	); erro != nil {
		return erro
	}

	return nil
}
