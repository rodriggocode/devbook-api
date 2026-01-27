package login

import (
	"devbook-api/app/auth"
	"devbook-api/app/database"
	"devbook-api/app/entity"
	repository "devbook-api/app/repository/login"
	"devbook-api/app/respostas"
	"devbook-api/app/secret"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// login para autenticar o usuario
func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user entity.UsersEntity
	if erro = json.Unmarshal(body, &user); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	respository := repository.NewRepositoryUser(db)
	userSaveDatabase, erro := respository.Login(user.Email)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = secret.CheckiPassword(userSaveDatabase.Password, user.Password); erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.GenerationToke(userSaveDatabase.ID)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
}
