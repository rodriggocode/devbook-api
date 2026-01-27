package posts

import (
	"api-web/app/auth"
	"api-web/app/database"
	"api-web/app/respostas"
	"net/http"
	repository "api-web/app/repository/posts"
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
