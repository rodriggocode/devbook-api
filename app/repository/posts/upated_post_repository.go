package posts

import (
	"database/sql"
	"devbook-api/app/entity"
)

type UpdatedRepository struct {
	db *sql.DB
}

func NewUpdatedRepository(db *sql.DB) UpdatedRepository {
	return UpdatedRepository{db}
}

func (repo UpdatedRepository) UpdatedPost(postID uint64, posts entity.Posts) error {
	statment, err := repo.db.Prepare("UPDATE posts SET title = $1, content = $2, WHERE id_post = $3")
	if err != nil {
		return err
	}

	defer statment.Close()

	if _, err = statment.Exec(posts.Title, posts.Content, postID); err != nil {
		return err
	}

	return nil
}
