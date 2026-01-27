package posts

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	repository "devbook-api/app/repository/posts"
	"devbook-api/app/respostas"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, req *http.Request) {
	userID, erro := auth.ExtractUserID(req)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewGetPostRepository(db)
	posts, erro := repository.GetAllPosts(userID)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)

}
