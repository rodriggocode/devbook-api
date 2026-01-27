package posts

import (
	"api-web/app/entity"
	"database/sql"
)

type GetAllPostsUserRepository struct {
	db *sql.DB
}

func NewGetAllPostsRepository(db *sql.DB) GetAllPostRepository {
	return GetAllPostRepository{db}
}

func (repo GetAllPostRepository) GetAllPostsUser(userID uint64) ([]entity.Posts, error) {
	rows, err := repo.db.Query(
		`SELECT p.id_posts, p.title, p.content, p.author_id, p.like_posts, p.created_at, u.nick
		FROM posts p
		JOIN users u ON u.id_user= p.author_id
		 WHERE p.author_id = $1`, userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []entity.Posts

	for rows.Next() {
		var post entity.Posts

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Like,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
