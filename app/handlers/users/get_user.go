package users

import (
	"api-web/app/database"
	repository "api-web/app/repository/users"
	"api-web/app/respostas"
	"net/http"
	"strings"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := database.Connection()
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewGetRepository(db)
	users, erro := repo.GetAll(nameOrNick)
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, users)
}
