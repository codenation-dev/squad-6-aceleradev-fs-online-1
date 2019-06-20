package models

import (
	"time"
)

// AlertHistory clientes do sistema
type AlertHistory struct {
	ID                int             `form:"id" json:"id"`
	Date              time.Time       `form:"date" json:"date"`
	User              User            `form:"user" json:"user"`
	Customer          Customer        `form:"customer" json:"customer"`
	PaymentEmployee   PaymentEmployee `form:"paymentEmployee" json:"paymentEmployee"`
	Payment           Payment         `form:"payment" json:"payment"`
	UserID            int             `form:"userid" json:"userId"`
	CustomerID        int             `form:"customerId" json:"customerId"`
	PaymentEmployeeID int             `form:"paymentEmployeeId" json:"paymentEmployeeId"`
}
