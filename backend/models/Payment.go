package models

// Payment modelo para pagamentos
type Payment struct {
	ID               int    `form:"id" json:"id" binding:"required"`
	FileName         string `form:"filename" json:"filename" binding:"required"`
	Month            int    `form:"month" json:"month" binding:"required"`
	Year             int    `form:"year" json:"year" binding:"required"`
	EmployeePayments []PaymentEmployee
}

// PaymentEmployee modelo para historico de pagamentos
type PaymentEmployee struct {
	ID         int      `form:"id" json:"id" binding:"required"`
	Name       string   `form:"name" json:"name" binding:"required"`
	Occupation string   `form:"occupation" json:"occupation" binding:"required"`
	Department string   `form:"department" json:"department" binding:"required"`
	Salary     float64  `form:"salary" json:"salary" binding:"required"`
	Customer   Customer `form:"customer" json:"customer" binding:"required"`
}
