package posts

import (
	"api-web/app/entity"
	"database/sql"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

func (repo PostsRepository) CreatePosts(posts entity.Posts) (uint64, error) {
	var postID uint64
	row := repo.db.QueryRow("INSERT INTO posts (title, content ,author_id, author_nick) VALUES($1, $2, $3, $4) RETURNING id_posts",
		posts.Title, posts.Content, posts.AuthorID, posts.AuthorNick)

	erro := row.Scan(&postID)
	if erro != nil {
		return 0, erro
	}

	return postID, nil
}
