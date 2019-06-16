package models

import (
	"time"
)

// AlertHistory clientes do sistema
type AlertHistory struct {
	ID              int             `form:"id" json:"id"`
	Date            time.Time       `form:"date" json:"date"`
	User            User            `form:"user" json:"user"`
	Customer        Customer        `form:"customer" json:"customer"`
	PaymentEmployee PaymentEmployee `form:"paymentEmployee" json:"paymentEmployee"`
}
