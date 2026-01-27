package posts

import (
	"api-web/app/auth"
	"api-web/app/database"
	"api-web/app/entity"
	repository "api-web/app/repository/posts"
	"api-web/app/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func UpdatedPost(w http.ResponseWriter, req *http.Request) {
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
	defer db.Close()

	searchRepo := repository.NewSearchIDPost(db)
	postFromDB, err := searchRepo.GetPostID(postID)
	if err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	if postFromDB.AuthorID != userID {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("nao e possivel atualizar uma publicacao que nao seja a sua"))
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		respostas.RespostaError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var postToUpdate entity.Posts
	if err = json.Unmarshal(body, &postToUpdate); err != nil {
		respostas.RespostaError(w, http.StatusBadRequest, err)
		return
	}

	updateRepo := repository.NewUpdatedPostRepository(db)
	if err = updateRepo.UpdatePost(postID, postToUpdate); err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
