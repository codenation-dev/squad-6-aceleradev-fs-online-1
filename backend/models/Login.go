package models

//Login modelo usado no login
type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
