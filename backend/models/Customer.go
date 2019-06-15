package models

// Customer clientes do sistema
type Customer struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}
