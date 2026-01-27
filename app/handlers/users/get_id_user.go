package users

import (
	"api-web/app/database"
	repository "api-web/app/repository/users"
	"api-web/app/respostas"
	"net/http"
	"strconv"
)

func GetIdUser(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id") // e esse id, eu pensava que tinha relacao com a tabel, ele e so o que vou passar na url: curl http://localhost:5000/usuarios_id?id=3

	if param == "" {
		http.Error(w, "Id nao informado", http.StatusBadRequest)
		return
	}
	//userID, erro := strconv.ParseUint(param, 10, 64) // aqui eu to convertendo um uint64, que ta la no tipo da minha entity
	userID, erro := strconv.Atoi(param)
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

	repo := repository.NewGetIDUser(db)
	user, erro := repo.GetID(uint64(userID))
	if erro != nil {
		respostas.RespostaError(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, user)
}
