package routes

import (
	handlers "devbook-api/app/handlers/posts"
	"devbook-api/app/middlewares"
	"net/http"
)

func LoadPostRoutes(router *http.ServeMux) {
	// Criar post
	router.HandleFunc("POST /publicacao",
		middlewares.Authentication(handlers.CreatePosts))

	// Feed
	router.HandleFunc("GET /publicacoes",
		middlewares.Authentication(handlers.GetAllPosts))

	// Buscar post por ID
	router.HandleFunc("GET /publicacao/{posts_id}",
		middlewares.Authentication(handlers.GetIDPost))

	// Editar
	router.HandleFunc("PUT /publicacao/{posts_id}",
		middlewares.Authentication(handlers.UpdatedPost))

	// Excluir
	router.HandleFunc("DELETE /publicacao/{posts_id}",
		middlewares.Authentication(handlers.DeletePost))

	// Posts de um usuário
	router.HandleFunc("GET /usuarios/{user_id}/publicacoes",
		middlewares.Authentication(handlers.GetAllPostsUser))

	// Curtir
	router.HandleFunc("POST /publicacoes/{posts_id}/curtir",
		middlewares.Authentication(handlers.LikePosts))

	// Descurtir
	router.HandleFunc("POST /publicacoes/{posts_id}/descurtir",
		middlewares.Authentication(handlers.Unliked))
}
