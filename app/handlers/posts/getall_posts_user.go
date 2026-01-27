package posts

import (
	"api-web/app/database"
	repository "api-web/app/repository/posts"
	"api-web/app/respostas"
	"net/http"
	"strconv"
)

func GetAllPostsUser(w http.ResponseWriter, req *http.Request) {
	params := req.PathValue("user_id")
	userID, err := strconv.ParseUint(params, 10, 64)
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

	repository := repository.NewGetAllPostsRepository(db)
	posts, err := repository.GetAllPostsUser(userID)

	if err != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, posts)
}
