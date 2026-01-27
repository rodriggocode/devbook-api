package entity

import (
	"errors"
	"strings"
	"time"
)

// aqui esse posts e relacionado a tabela de publicacoes dos ususarios

type Posts struct {
	ID         uint64    `json:"id_posts,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Like       uint64    `json:"like"`
	CreatedAt  time.Time `json:"created_at"`
}

func (post *Posts) Prepare() error {
	if erro := post.validate(); erro != nil {
		return erro
	}

	post.format()
	return nil
}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("O titulo e obrigatorio e nao pode ser em branco")
	}
	if post.Content == "" {
		return errors.New("O conteudo nao pode ser em branco")
	}

	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
