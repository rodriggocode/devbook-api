package users

import (
	"devbook-api/app/database"
	"devbook-api/app/entity"
	repository "devbook-api/app/repository/users"
	"devbook-api/app/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var user entity.UsersEntity
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}
	// esse parametro, ele vem por conta de umas mudancas nos metodos la na entity
	if erro = user.Preparar("create"); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, erro = repository.CreateUser(user)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, user)
}
