package login

import (
	"api-web/app/entity"
	"database/sql"
)

type LoginRepository struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) LoginRepository {
	return LoginRepository{
		db: db,
	}
}

func (repositroy LoginRepository) Login(email string) (entity.UsersEntity, error) {
	row, erro := repositroy.db.Query("SELECT id_user, password FROM users WHERE email = $1", email)
	if erro != nil {
		return entity.UsersEntity{}, erro
	}
	defer row.Close()

	var user entity.UsersEntity

	if row.Next() {
		if erro = row.Scan(
			&user.ID,
			&user.Password,
		); erro != nil {
			return entity.UsersEntity{}, erro
		}
	}
	return user, nil
}
