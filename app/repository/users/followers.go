package users

import (
	"api-web/app/entity"
	"database/sql"
)

type FollowersRepository struct {
	db *sql.DB
}

func NewFollow(db *sql.DB) *FollowersRepository {
	return &FollowersRepository{db}
}

// aqui eu vou colocar o id para aprecer, so no modo dev, depois eu tiro
func (repo FollowersRepository) Follow(userID uint64) ([]entity.UsersEntity, error) {
	rows, erro := repo.db.Query(`
    	SELECT u.id_user, u.user_name, u.nick, u.email, u.created_at
    	FROM users u INNER JOIN follows f  ON u.id_user = user_id
    	WHERE f.follower_id = $1
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
			&user.Nick,
			&user.Email,
			&user.Created_at,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil

}
