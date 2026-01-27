package users

import (
	"database/sql"
)

type FollowUserRepository struct {
	db *sql.DB
}

func NewRepositoryFollowUser(db *sql.DB) *FollowUserRepository {
	// aqui poderia colocar so o db, mas eu quis fazer diferente, para saber que
	// tem mais de uma maneira de fazer
	return &FollowUserRepository{db: db}
}

func (repo FollowUserRepository) FollowUser(followID, userID uint64) error {
	statement, erro := repo.db.Prepare("INSERT INTO follows(user_id, follower_id) VALUES ($1, $2)")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(followID, userID); erro != nil {
		return erro
	}

	return nil

}
