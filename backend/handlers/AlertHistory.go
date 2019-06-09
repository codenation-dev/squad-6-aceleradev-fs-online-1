package handlers

import (
	"os"
	"strconv"
	"time"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/services"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
)

//RegisterAndNotifyAlerts a
func RegisterAndNotifyAlerts(paymentID int) {

	var listAlert []models.AlertHistory
	var listUserAlert []models.User

	minSalaryForRegisterPayment, err := strconv.ParseFloat(os.Getenv("CONFIG_MIN_SALARY_REGISTER_PAYMENT"), 64)
	if err != nil {
		minSalaryForRegisterPayment = 20000
	}
	payment := db.FindPaymentByID(true, paymentID)

	for _, employee := range payment.EmployeePayments {
		if employee.Customer.ID > 0 && employee.Salary >= minSalaryForRegisterPayment {
			listUserAlert = db.FindAllUsersReceiveAlert()
			for _, user := range listUserAlert {

				alertHistory := models.AlertHistory{
					Customer:        employee.Customer,
					Date:            time.Now(),
					PaymentEmployee: employee,
					User:            user,
				}
				alertInserted := db.InsertAlertHistory(alertHistory)
				listAlert = append(listAlert, alertInserted)
			}
		}
	}

	if len(listAlert) > 0 && len(listAlert) > 0 {
		services.SendEmailAlertEmployeeSalary(listUserAlert, listAlert)
	}

}
