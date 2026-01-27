package users

import (
	"api-web/app/auth"
	"api-web/app/database"
	repository "api-web/app/repository/users"
	"api-web/app/respostas"
	"errors"
	"net/http"
	"strconv"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	params := r.PathValue("usuarios_id")
	userID, erro := strconv.Atoi(params) // teste, saber se vai funfar
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}
	if followID == uint64(userID) {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Nao e possivel seguir voce mesmo"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryFollowUser(db)
	if erro = repository.FollowUser(followID, uint64(userID)); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
