package users

import (
	"database/sql"
)

type NovaSenhaRepository struct {
	db *sql.DB
}

func AtualizarSenha(db *sql.DB) NovaSenhaRepository {
	return NovaSenhaRepository{db}
}

func (repo NovaSenhaRepository) AtulizarSenhaDB(userID uint64, password string) error {
	statement, erro := repo.db.Prepare("UPDATE users SET password=$1 WHERE id_user=$2")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}
