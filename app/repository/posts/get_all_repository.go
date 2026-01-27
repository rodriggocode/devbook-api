package posts

import (
	"database/sql"
	"devbook-api/app/entity"
)

type GetAllPostRepository struct {
	db *sql.DB
}

func NewGetPostRepository(db *sql.DB) GetAllPostRepository {
	return GetAllPostRepository{db: db}
}

func (repo *GetAllPostRepository) GetAllPosts(userID uint64) ([]entity.Posts, error) {
	rows, erro := repo.db.Query(`
		SELECT DISTINCT p.id_posts, p.title, p.content, p.author_id, p.like_posts, p.created_at, u.nick
		FROM posts p
		INNER JOIN users u ON u.id_user = p.author_id
		INNER JOIN follows f ON p.author_id = f.user_id
		WHERE u.id_user = $1 OR f.follower_id = $1 ORDER BY 1`,
		userID,
	)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var posts []entity.Posts

	for rows.Next() {
		var post entity.Posts
		if erro = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Like,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil

}
