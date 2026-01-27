package posts

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	repostiory "devbook-api/app/repository/posts"
	"devbook-api/app/respostas"
	"errors"
	"net/http"
	"strconv"
)

func DeletePost(w http.ResponseWriter, req *http.Request) {
	userID, err := auth.ExtractUserID(req)
	if err != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, err)
		return
	}

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

	searcRepo := repostiory.NewSearchIDPost(db)
	postSave, err := searcRepo.GetPostID(uint64(postID))

	if err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	if postSave.AuthorID != userID {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Voce nao pode excluir uma publicacao que nao seja sua!"))
	}

	delete := repostiory.NewDeletePostRepository(db)
	if err = delete.DeletePost(postID); err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
