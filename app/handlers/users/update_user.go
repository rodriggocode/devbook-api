package users

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	"devbook-api/app/entity"
	repository "devbook-api/app/repository/users"
	"devbook-api/app/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paramID := r.URL.Query().Get("id")
	userId, erro := strconv.Atoi(paramID)

	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
	}

	if uint64(userId) != userIDToken {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Nao e possivel atualizar um usuario que nao seja o seu"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var users entity.UsersEntity
	if erro = json.Unmarshal(body, &users); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	// tem esse parametro, por conta de umas mudancas la na entity
	if erro = users.Preparar("edit"); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewUpatedRepository(db)

	if erro = repository.UpatedUser(uint64(userId), users); erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, repository)

}
