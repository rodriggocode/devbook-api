package users

import (
	"api-web/app/database"
	repository "api-web/app/repository/users"
	"api-web/app/respostas"
	"net/http"
	"strconv"
)

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("usuarios_id")
	userID, erro := strconv.Atoi(params)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	repository := repository.NewGetFollowRepository(db)
	user_ID, erro := repository.GetFollowers(uint64(userID))
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, user_ID)

}
