package models

// Payment modelo para pagamentos
type Payment struct {
	ID               int    `form:"id" json:"id"`
	FileName         string `form:"filename" json:"filename"`
	Month            int    `form:"month" json:"month"`
	Year             int    `form:"year" json:"year"`
	EmployeePayments []PaymentEmployee
}

// PaymentEmployee modelo para historico de pagamentos
type PaymentEmployee struct {
	ID         int      `form:"id" json:"id"`
	Name       string   `form:"name" json:"name"`
	Occupation string   `form:"occupation" json:"occupation"`
	Department string   `form:"department" json:"department"`
	Salary     float64  `form:"salary" json:"salary"`
	Customer   Customer `form:"customer" json:"customer"`
}
