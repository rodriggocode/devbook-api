package posts

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	"devbook-api/app/entity"
	repository "devbook-api/app/repository/posts"
	"devbook-api/app/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var posts entity.Posts
	if erro = json.Unmarshal(bodyRequest, &posts); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	posts.AuthorID = userID

	if erro = posts.Prepare(); erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repository.NewPostRepository(db)
	posts.ID, erro = repository.CreatePosts(posts)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, posts)

}
