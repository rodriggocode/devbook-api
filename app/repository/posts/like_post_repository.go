package posts

import (
	"database/sql"
	"errors"
)

type LikePostsRepo struct {
	db *sql.DB
}

func NewLikePostsRepository(db *sql.DB) LikePostsRepo {
	return LikePostsRepo{db}
}

func (rep LikePostsRepo) LikePosts(postID uint64) error {
	statement, err := rep.db.Prepare("UPDATE posts SET like_posts = like_posts + 1 WHERE id_posts = $1 RETURNING id_posts")
	if err != nil {
		return err
	}
	defer statement.Close()

	var updatedID uint64
	err = statement.QueryRow(postID).Scan(&updatedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("post não encontrado ou nenhuma linha afetada")
		}
		return err
	}

	return nil
}
