package posts

import (
	"devbook-api/app/database"
	repository "devbook-api/app/repository/posts"
	"devbook-api/app/respostas"
	"net/http"
	"strconv"
)

func GetIDPost(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("posts_id")
	postID, erro := strconv.Atoi(params)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadGateway, erro)
		return
	}

	db, erro := database.Connection()

	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewSearchIDPost(db)
	post, erro := repository.GetPostID(uint64(postID))
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, post)
}
