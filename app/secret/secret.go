package secret

import "golang.org/x/crypto/bcrypt"

// Has recebe uma string e coloca em um has nela
func Has(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// verficar senha e compara uma senha e um has e retorna se sao iguais
func CheckiPassword(passWordWithHas, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passWordWithHas), []byte(passwordString))
}
