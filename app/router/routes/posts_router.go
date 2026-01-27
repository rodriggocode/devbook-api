package routes

import (
	handlers "devbook-api/app/handlers/posts"
	"devbook-api/app/middlewares"
	"net/http"
)

func LoadPostRoutes(router *http.ServeMux) {
	router.HandleFunc("/publicacao/criar", middlewares.Authentication(handlers.CreatePosts))
	router.HandleFunc("/publicacoes", middlewares.Authentication(handlers.GetAllPosts))
	router.HandleFunc("/publicacao/{posts_id}/publicacao", middlewares.Authentication(handlers.GetIDPost))
	router.HandleFunc("/publicacao/{posts_id}/editar", middlewares.Authentication(handlers.UpdatedPost))
	router.HandleFunc("/publicacao/{posts_id}/excluir", middlewares.Authentication(handlers.DeletePost))

	router.HandleFunc("/publicacao/{user_id}/usuarios", middlewares.Authentication(handlers.GetAllPostsUser))
	router.HandleFunc("/publicacoes/{posts_id}/curtir", middlewares.Authentication(handlers.LikePosts))
	router.HandleFunc("/publicacao/{posts_id}/descurtir", middlewares.Authentication(handlers.Unliked))
}
