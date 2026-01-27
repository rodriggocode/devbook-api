package auth

import (
	"devbook-api/app/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// go get github.com/dgrijalva/jwt-go

func GenerationToke(userID uint64) (string, error) {
	// definir permissoes
	permission := jwt.MapClaims{}
	permission["authorized"] = true                          // quem tiver o toke ele vai autorizar
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix() // tempo que o token vai expirar, o 6 que dizer que vai levar 6h para expirar e o uni
	permission["id_user"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	return token.SignedString([]byte(config.SecretKey)) //secret
}

// aqui vai validar o se o token e valido mesmo que ta sendo passado na
// requisicao
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token Invalido")

}

// extrair um ususarioid e retorna um usuario que ta
// salvo no token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKey)
	if erro != nil {
		return 0, erro
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permission["id_user"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userID, nil
	}

	return 0, errors.New("Token Invalido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inseperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
