package models

import (
	"time"
)

// AlertHistory clientes do sistema
type AlertHistory struct {
	ID              int             `form:"id" json:"id" binding:"required"`
	Date            time.Time       `form:"date" json:"date" binding:"required"`
	User            User            `form:"user" json:"user" binding:"required"`
	Customer        Customer        `form:"customer" json:"customer" binding:"required"`
	PaymentEmployee PaymentEmployee `form:"paymentEmployee" json:"paymentEmployee" binding:"required"`
}
