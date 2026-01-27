package posts

import "database/sql"

type DeletePost struct {
	db *sql.DB
}

func NewDeletePostRepository(db *sql.DB) DeletePost {
	return DeletePost{db}
}

func (repo DeletePost) DeletePost(postID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM posts WHERE id_posts = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
