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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("id")
	userId, erro := strconv.Atoi(params)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	userToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	if uint64(userId) != userToken {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Voce nao tem permissao para deletar um usuario que nao seja o seu"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}
	defer db.Close()

	repository := repository.NewDeleteUser(db)
	// esse(uint64(variavel)) e opcional, vou deixar para testar se vai da certo
	if erro = repository.DeletUser(uint64(userId)); erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
