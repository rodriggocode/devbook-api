package login

import (
	"devbook-api/app/auth"
	cookie "devbook-api/app/config"
	"devbook-api/app/database"
	"devbook-api/app/entity"
	repository "devbook-api/app/repository/login"
	"devbook-api/app/respostas"
	"devbook-api/app/secret"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://webapp-snowy-flower-2545.fly.dev")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Vary", "Origin") 
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://webapp-snowy-flower-2545.fly.dev")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Vary", "Origin") 
	body, erro := io.ReadAll(r.Body)
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

	userID := strconv.FormatUint(userSaveDatabase.ID, 10)

	cookie.SetCookie(w, token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.DateAuth{
		ID:   userID,
		Toke: token,
	})

}
