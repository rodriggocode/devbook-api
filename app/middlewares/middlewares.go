package middlewares

import (
	"devbook-api/app/auth"
	"devbook-api/app/respostas"
	"log"
	"net/http"
)

// aqui so vai escrever as informacoes das requisicoes no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// autenticar e verifica se o usuario que ta fazendo a autenticacao esta
// autenticado de fato!
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			respostas.RespostaError(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
