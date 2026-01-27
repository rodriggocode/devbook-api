package posts

import (
	"api-web/app/database"
	repository "api-web/app/repository/posts"
	"api-web/app/respostas"
	"net/http"
	"strconv"
)

func Unliked(w http.ResponseWriter, req *http.Request) {
	params := req.PathValue("posts_id")
	postID, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		respostas.RespostaError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()

	if err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUnlikedRepository(db)
	if err = repository.Unliked(postID); err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
