package users

import (
	"api-web/app/entity"
	"database/sql"
)

type SearchPasswordRepository struct {
	db *sql.DB
}

func NewPasswordUser(db *sql.DB) SearchPasswordRepository {
	return SearchPasswordRepository{db}

}

func (repo SearchPasswordRepository) SearchPassword(userId uint64) (string, error) {
	row, erro := repo.db.Query("SELECT password FROM users WHERE id_user=$1", userId)
	if erro != nil {
		return "", erro
	}
	defer row.Close()

	var user entity.UsersEntity
	if row.Next() {
		if erro = row.Scan(
			&user.Password,
		); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}
