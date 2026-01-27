package users

import "database/sql"

type DeleteUserRepository struct {
	db *sql.DB
}

func NewDeleteUser(db *sql.DB) DeleteUserRepository {
	return DeleteUserRepository{db}
}

func (repo DeleteUserRepository) DeletUser(ID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM users WHERE id_user = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}
