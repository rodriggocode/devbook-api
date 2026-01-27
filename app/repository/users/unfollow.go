package users

import (
	"database/sql"
)

type UnFollowRepository struct {
	db *sql.DB
}

func NewUnFollowRepository(db *sql.DB) *UnFollowRepository {
	return &UnFollowRepository{db}
}

func (repo UnFollowRepository) UnFollow(followID, userID uint64) error {
	statment, erro := repo.db.Prepare("DELETE FROM follows WHERE user_id = $1 AND follower_id = $2")
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(followID, userID); erro != nil {
		return erro
	}

	return nil
}
