package entity

// aqui vai servir para alteracao de senha
type PasswordEntity struct {
	NewPassword     string `json:"new_password"`
	CurrentPassword string `json:"current_password"`
}
