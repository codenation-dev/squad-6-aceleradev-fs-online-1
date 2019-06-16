package handlers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/services"
	"github.com/gin-gonic/gin"

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

// GetAlerts retorna todos alertas, busca por usuario, cliente e pagamento
func GetAlerts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	userID, _ := strconv.Atoi(c.Query("userId"))
	customerID, _ := strconv.Atoi(c.Query("customerId"))
	paymentID, _ := strconv.Atoi(c.Query("paymentId"))

	list := db.FindAlerts(userID, customerID, paymentID, 0)

	if len(list) > 0 {
		c.JSON(http.StatusOK, list)
	} else {
		c.JSON(http.StatusNoContent, nil)
	}

}

// GetAlert retorna alerta por id
func GetAlert(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var alert models.AlertHistory

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {

		listAlert := db.FindAlerts(0, 0, 0, id)

		if len(listAlert) > 0 {
			alert = listAlert[0]
		}

		if alert.ID > 0 {
			c.JSON(http.StatusOK, alert)
		} else {
			c.AbortWithStatus(http.StatusNoContent)
		}

	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
