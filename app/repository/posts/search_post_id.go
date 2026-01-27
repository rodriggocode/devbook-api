package posts

import (
	"api-web/app/entity"
	"database/sql"
)

type SearchIDPost struct {
	db *sql.DB
}

func NewSearchIDPost(db *sql.DB) *SearchIDPost {
	return &SearchIDPost{db}
}

func (repo *SearchIDPost) GetPostID(postID uint64) (entity.Posts, error) {
	row, erro := repo.db.Query(`
		SELECT p.id_posts, p.title, p.content, p.author_id, u.nick, p.like_posts, p.created_at FROM
		posts p INNER JOIN users u
		ON u.id_user = p.author_id
		WHERE p.id_posts = $1
	`, postID)

	if erro != nil {
		return entity.Posts{}, erro
	}

	defer row.Close()

	var post entity.Posts

	if row.Next() {
		if erro = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorNick,
			&post.Like,
			&post.CreatedAt,
		); erro != nil {
			return entity.Posts{}, erro
		}
	}

	return post, nil
}
