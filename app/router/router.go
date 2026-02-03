package router

import (
	"devbook-api/app/handlers/login"
	handlers "devbook-api/app/handlers/users"
	"devbook-api/app/middlewares"
	"devbook-api/app/router/routes"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/usuarios", middlewares.Authentication(handlers.GetUser))
	router.HandleFunc("/create/user", handlers.CreateUser)
	router.HandleFunc("/usuarios_id", middlewares.Authentication(handlers.GetIdUser))
	router.HandleFunc("/usuarios/id_editar", middlewares.Authentication(handlers.UpdateUser))
	router.HandleFunc("/usuario/excluir", middlewares.Authentication(handlers.DeleteUser))

	// login, tem que ser um metodo post
	router.HandleFunc("/login", login.Login)

	// aqui tambem vai ser um metodo post
	router.HandleFunc("/usuarios/{usuarios_id}/seguir", middlewares.Authentication(handlers.FollowUser))
	router.HandleFunc("/usuarios/{usuarios_id}/deixar-de-seguir", middlewares.Authentication(handlers.UnFollow))
	router.HandleFunc("/usuarios/{usuarios_id}/seguidores", middlewares.Authentication(handlers.GetFollowers))
	router.HandleFunc("/usuarios/{usuarios_id}/seguindo", middlewares.Authentication(handlers.Follow))
	router.HandleFunc("/usuarios/{usuarios_id}/atualizar-senha", middlewares.Authentication(handlers.UpadatedPassword))

	// aqui teve uma mudanca, eu estou separando as rotas
	// para que nao fique tudo centralizado em um unico arquivo
	routes.LoadPostRoutes(router)

	return router
}
