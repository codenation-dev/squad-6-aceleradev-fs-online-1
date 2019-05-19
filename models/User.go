package models

// User usuario do sistema
type User struct {
	ID           int    `form:"id" json:"id" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Name         string `form:"name" json:"name" binding:"required"`
	ReceiveAlert bool   `form:"receiveAlert" json:"receiveAlert" binding:"required"`
}
