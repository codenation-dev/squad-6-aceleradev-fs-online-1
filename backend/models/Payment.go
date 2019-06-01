package models

// Payment modelo para pagamentos
type Payment struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	FileName string `form:"filename" json:"filename" binding:"required"`
	Month    int    `form:"month" json:"month" binding:"required"`
	Year     int    `form:"year" json:"year" binding:"required"`
}

type PaymentHistory struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	FileName string `form:"filename" json:"filename" binding:"required"`
	Month    int    `form:"month" json:"month" binding:"required"`
	Year     int    `form:"year" json:"year" binding:"required"`
}