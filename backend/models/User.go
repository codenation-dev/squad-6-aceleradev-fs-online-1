package models

// User usuario do sistema
type User struct {
	ID           int    `form:"id" json:"id"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
	Name         string `form:"name" json:"name"`
	ReceiveAlert bool   `form:"receiveAlert" json:"receiveAlert"`
}
