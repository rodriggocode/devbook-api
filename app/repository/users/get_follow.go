package users

import (
	"api-web/app/entity"
	"database/sql"
)

type GetFollowRepository struct {
	db *sql.DB
}

func NewGetFollowRepository(db *sql.DB) GetFollowRepository {
	return GetFollowRepository{db}
}

func (repo GetFollowRepository) GetFollowers(userID uint64) ([]entity.UsersEntity, error) {
	rows, erro := repo.db.Query(`
    	SELECT u.id_user,  u.user_name,u.nick , u.email, u.created_at, u.updated_at
    	FROM users u INNER JOIN follows f on u.id_user = f.follower_id WHERE f.user_id = $1
    	`, userID)

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
			&user.Email,
			&user.Nick,
			&user.Created_at,
			&user.Updated_at,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}
