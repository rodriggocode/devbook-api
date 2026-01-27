package users

import (
	"api-web/app/auth"
	"api-web/app/database"
	"api-web/app/entity"
	rep "api-web/app/repository/users"
	repository "api-web/app/repository/users"
	"api-web/app/respostas"
	"api-web/app/secret"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func UpadatedPassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, erro)
		return
	}

	params := r.PathValue("usuarios_id") // aqui eu acho que nao e o r. mas sim userIDToken
	userID, erro := strconv.Atoi(params)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadGateway, erro)
		return
	}

	if userIDToken != uint64(userID) {
		respostas.RespostaError(w, http.StatusForbidden, errors.New("Nao e possivel atualizar a senha de  um usuario que nao seja o seu"))
		return
	}

	// Lê todo o conteúdo do corpo da requisição.
	// io.ReadAll substitui o depreciado ioutil.ReadAll.
	bodyRequest, erro := io.ReadAll(r.Body)

	var password entity.PasswordEntity
	if erro = json.Unmarshal(bodyRequest, &password); erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewPasswordUser(db)
	newPasswordDataba, erro := repository.SearchPassword(uint64(userID))

	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = secret.CheckiPassword(newPasswordDataba, password.CurrentPassword); erro != nil {
		respostas.RespostaError(w, http.StatusUnauthorized, errors.New("A senha atual no e igual a do banco de dados"))
		return
	}

	HasPassword, erro := secret.Has(password.NewPassword)
	if erro != nil {
		respostas.RespostaError(w, http.StatusBadRequest, erro)
		return
	}

	// aqui ficou como AtualizarSenha, pq ja tem um repository de UpadatedPassword
	repoAtualizarSenha := rep.AtualizarSenha(db)
	if erro = repoAtualizarSenha.AtulizarSenhaDB(uint64(userID), string(HasPassword)); erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
