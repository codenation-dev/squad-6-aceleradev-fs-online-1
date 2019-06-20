package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
)

//InsertAlertHistory cadastra historico de alerta e retorna registro inserido no banco
func InsertAlertHistory(optionalDB *sql.DB, alertHistory models.AlertHistory) models.AlertHistory {
	var (
		ID                   int
		alertHistoryInserted models.AlertHistory
	)

	var db *sql.DB
	if optionalDB != nil {
		db = optionalDB
	} else {
		db = ConnectDataBase()
		defer CloseDataBase(db)
	}

	insert :=
		`INSERT INTO public.historico_alerta
		(hisale_data, usuari_id, client_id, pagfun_id)
		VALUES ($1, $2, NULLIF($3,0), $4) 
		returning hisale_id;`

	errUpdate := db.QueryRow(insert,
		alertHistory.Date,
		alertHistory.User.ID,
		alertHistory.Customer.ID,
		alertHistory.PaymentEmployee.ID).Scan(&ID)

	if errUpdate != nil {
		log.Println("db.InsertAlertHistory->Erro ao executar insert. Error:", errUpdate)
	} else {
		alertHistoryInserted = alertHistory
		alertHistoryInserted.ID = ID
	}
	return alertHistoryInserted
}

//FindAlerts retorna alertas
func FindAlerts(userID int, customerID int, paymentID int, ID int, onlyCustomers int) []models.AlertHistory {
	var (
		alertID                        int
		alertDate                      time.Time
		alertUserID                    int
		alertUserName                  string
		alertUserEmail                 string
		alertCustomerID                int
		alertCustomerName              string
		alertPaymentEmployeeID         int
		alertPaymentEmployeeName       string
		alertPaymentEmployeeOccupation string
		alertPaymentEmployeeDepartment string
		alertPaymentEmployeeSalary     float64
		alertPaymentID                 int
		alertPaymentYear               int
		alertPaymentMonth              int
		list                           []models.AlertHistory
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	sql := `select 
				historico_alerta.hisale_id,
				historico_alerta.hisale_data,
				historico_alerta.usuari_id,
				coalesce(historico_alerta.client_id, 0) as client_id,
				historico_alerta.pagfun_id,	
				usuario.usuari_nome,
				usuario.usuari_email,
				coalesce(cliente.client_nome, '') as client_nome,
				pagamento.pagame_id,
				pagamento.pagame_ano,
				pagamento.pagame_mes,
				pagamento_funcionario.pagfun_nome,
				pagamento_funcionario.pagfun_cargo,
				pagamento_funcionario.pagfun_orgao,
				pagamento_funcionario.pagfun_remuneracao
			from historico_alerta 			

			left join pagamento_funcionario on 
				(pagamento_funcionario.pagfun_id = historico_alerta.pagfun_id)

			left join pagamento on 
				(pagamento_funcionario.pagame_id = pagamento.pagame_id)				

			left join cliente on 
				(cliente.client_id = historico_alerta.client_id)

			left join usuario on 
				(usuario.usuari_id = historico_alerta.usuari_id)
			
			`

	sql = sql + "where (true) "
	if ID > 0 {
		sql = sql + ` and (hisale_id =` + strconv.Itoa(ID) + `)`
	}
	if paymentID > 0 {
		sql = sql + ` and (pagamento_funcionario.pagame_id = ` + strconv.Itoa(paymentID) + `)`
	}
	if customerID > 0 {
		sql = sql + ` and (cliente.client_id = ` + strconv.Itoa(customerID) + `)`
	}
	if userID > 0 {
		sql = sql + ` and (usuario.usuari_id = ` + strconv.Itoa(userID) + `)`
	}
	if onlyCustomers == 1 {
		sql = sql + ` and (cliente.client_id is not null)`
	}

	rows, errQuery := db.Query(sql)
	if errQuery != nil {
		log.Println("db.FindAlerts()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {
		err := rows.Scan(&alertID, &alertDate,
			&alertUserID, &alertCustomerID, &alertPaymentEmployeeID,
			&alertUserName, &alertUserEmail, &alertCustomerName,
			&alertPaymentID, &alertPaymentYear, &alertPaymentMonth,
			&alertPaymentEmployeeName, &alertPaymentEmployeeOccupation, &alertPaymentEmployeeDepartment,
			&alertPaymentEmployeeSalary,
		)
		if err != nil {
			log.Println("db.FindAlerts()->Erro ao executar consulta. Error:", err)
		} else {
			var payment = models.AlertHistory{
				ID:         alertID,
				Date:       alertDate,
				CustomerID: alertCustomerID,
				UserID:     alertUserID,
				Payment: models.Payment{
					ID:    alertPaymentID,
					Year:  alertPaymentYear,
					Month: alertPaymentMonth,
				},
				PaymentEmployee: models.PaymentEmployee{
					ID:         alertPaymentEmployeeID,
					Name:       alertPaymentEmployeeName,
					Occupation: alertPaymentEmployeeOccupation,
					Department: alertPaymentEmployeeDepartment,
					Salary:     alertPaymentEmployeeSalary,
				},
				User: models.User{
					ID:    alertUserID,
					Name:  alertUserName,
					Email: alertUserEmail,
				},
				Customer: models.Customer{
					ID:   alertCustomerID,
					Name: alertCustomerName,
				},
			}

			list = append(list, payment)
		}

	}

	return list
}
