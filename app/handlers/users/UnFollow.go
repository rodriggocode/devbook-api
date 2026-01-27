package users

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	repository "devbook-api/app/repository/users"
	"devbook-api/app/respostas"
	"errors"
	"net/http"
	"strconv"
)

func UnFollow(w http.ResponseWriter, r *http.Request) {
	followID, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	params := r.PathValue("usuarios_id")
	userID, erro := strconv.Atoi(params)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()

	if followID == uint64(userID) {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Voce nao pode deixar de se seguir"))
		return
	}

	defer db.Close()

	repository := repository.NewUnFollowRepository(db)
	if erro = repository.UnFollow(followID, uint64(userID)); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	respostas.RespostaError(w, http.StatusNoContent, nil)
}
