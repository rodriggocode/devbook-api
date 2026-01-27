package posts

import (
	"api-web/app/entity"
	"database/sql"
)

// UpdatedPostRepository é um repositório para atualizar posts
type UpdatedPostRepository struct {
	db *sql.DB
}

// NewUpdatedPostRepository cria um novo repositório para atualizar posts
func NewUpdatedPostRepository(db *sql.DB) *UpdatedPostRepository {
	return &UpdatedPostRepository{db}
}

// UpdatePost atualiza um post no banco de dados
func (repo *UpdatedPostRepository) UpdatePost(postID uint64, post entity.Posts) error {
	statement, erro := repo.db.Prepare("update posts set title = $1, content = $2 where id_posts = $3")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postID); erro != nil {
		return erro
	}

	return nil
}
