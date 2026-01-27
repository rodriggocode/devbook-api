package posts

import "database/sql"

type UnlikedRepo struct {
	db *sql.DB
}

func NewUnlikedRepository(db *sql.DB) UnlikedRepo {
	return UnlikedRepo{db}
}

func (repo UnlikedRepo) Unliked(postID uint64) error {
	statement, err := repo.db.Prepare(`
			UPDATE posts SET like_posts =
			CASE
				WHEN like_posts > 0 THEN like_posts - 1
			ELSE 0 END WHERE id_posts = $1
	`)

	if err != nil {
		return err
	}

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
