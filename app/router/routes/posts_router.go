package routes

import (
	handlers "devbook-api/app/handlers/posts"
	"devbook-api/app/middlewares"
	"net/http"
)

func LoadPostRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /publicacao/criar", middlewares.Authentication(handlers.CreatePosts))
	router.HandleFunc("GET /publicacoe", middlewares.Authentication(handlers.GetAllPosts))
	router.HandleFunc("GET /publicacao/{posts_id}/publicacao", middlewares.Authentication(handlers.GetIDPost))
	router.HandleFunc("PUT /publicacao/{posts_id}/editar", middlewares.Authentication(handlers.UpdatedPost))
	router.HandleFunc("DELETE /publicacao/{posts_id}/excluir", middlewares.Authentication(handlers.DeletePost))

	router.HandleFunc("GET /publicacao/{user_id}/usuarios", middlewares.Authentication(handlers.GetAllPostsUser))
	router.HandleFunc("GET /publicacoes/{posts_id}/curtir", middlewares.Authentication(handlers.LikePosts))
	router.HandleFunc("POST /publicacao/{posts_id}/descurtir", middlewares.Authentication(handlers.Unliked))
}
