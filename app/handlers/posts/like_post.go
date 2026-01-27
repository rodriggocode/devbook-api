package posts

import (
	"devbook-api/app/database"
	repository "devbook-api/app/repository/posts"
	"devbook-api/app/respostas"
	"net/http"
	"strconv"
)

func LikePosts(w http.ResponseWriter, req *http.Request) {
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

	repository := repository.NewLikePostsRepository(db)
	if err := repository.LikePosts(postID); err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
