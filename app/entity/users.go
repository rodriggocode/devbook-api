package entity

import (
	"devbook-api/app/secret"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type UsersEntity struct {
	ID         uint64    `json:"id_user,omitempty"`
	Nome       string    `json:"user_name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

// aqui vai receber um etapa, para eu saner se e um
// etapa de create or updated
func (users *UsersEntity) Preparar(etapa string) error {
	if erro := users.validar(etapa); erro != nil {
		return erro
	}
	if erro := users.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (users *UsersEntity) validar(etapa string) error {
	if users.Nome == "" {
		return errors.New("O campo nome nao pode ser vazio")
	}

	if users.Nick == "" {
		return errors.New("O nick nao pode ser um campo em branco, muito menos em preto")
	}

	if len(users.Nick) < 6 {
		return errors.New("O nick nao pode ser menor que 6 caracteres ")
	}

	if users.Email == "" {
		return errors.New("O campo email nao pode ficar vazio, e tem que ser um email valido")
	}

	if erro := checkmail.ValidateFormat(users.Email); erro != nil {
		return errors.New("O email tem que ser um tipo valido")
	}

	if etapa == "create" && users.Password == "" {
		return errors.New("Senha obrigatoria")
	}

	if etapa == "create" && len(users.Password) < 6 {
		return errors.New("A senha tem que ter mais de 6 digitos")
	}

	return nil
}

func (users *UsersEntity) formatar(etapa string) error {
	users.Nome = strings.TrimSpace(users.Nome)
	users.Nick = strings.TrimSpace(users.Nick)
	users.Email = strings.TrimSpace(users.Email)
	users.Password = strings.TrimSpace(users.Password)

	if etapa == "create" {
		passwordHas, erro := secret.Has(users.Password)
		if erro != nil {
			return erro
		}
		users.Password = string(passwordHas)
	}

	return nil
}
