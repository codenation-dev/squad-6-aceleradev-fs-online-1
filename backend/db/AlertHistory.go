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
func FindAlerts(userID int, customerID int, paymentID int, ID int) []models.AlertHistory {
	var (
		alertID                int
		alertDate              time.Time
		alertUserID            int
		alertCustomerID        int
		alertPaymentEmployeeID int
		list                   []models.AlertHistory
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	sql := `select 
				historico_alerta.hisale_id,
				historico_alerta.hisale_data,
				historico_alerta.usuari_id,
				coalesce(historico_alerta.client_id, 0) as client_id,
				historico_alerta.pagfun_id				
			from historico_alerta `

	if paymentID > 0 {
		sql = sql + " " +
			`inner join pagamento_funcionario on 
				(pagamento_funcionario.pagfun_id = historico_alerta.pagfun_id) and
				(pagamento_funcionario.pagame_id = ` + strconv.Itoa(paymentID) + `)`
	}
	if customerID > 0 {
		sql = sql + " " +
			`inner join cliente on 
				(cliente.client_id = historico_alerta.client_id) and
				(cliente.client_id = ` + strconv.Itoa(customerID) + `)`
	}
	if userID > 0 {
		sql = sql + " " +
			`inner join usuario on 
				(usuario.usuari_id = historico_alerta.usuari_id) and
				(usuario.usuari_id = ` + strconv.Itoa(userID) + `)`
	}

	if ID > 0 {
		sql = sql + " " +
			`where (hisale_id =` + strconv.Itoa(ID) + `)`
	}

	rows, errQuery := db.Query(sql)
	if errQuery != nil {
		log.Println("db.FindAlerts()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {
		err := rows.Scan(&alertID, &alertDate, &alertUserID, &alertCustomerID, &alertPaymentEmployeeID)
		if err != nil {
			log.Println("db.FindAlerts()->Erro ao executar consulta. Error:", err)
		} else {
			var payment = models.AlertHistory{
				ID:                alertID,
				Date:              alertDate,
				CustomerID:        alertCustomerID,
				UserID:            alertUserID,
				PaymentEmployeeID: alertPaymentEmployeeID,
				User:              FindUserByID(alertUserID),
				Customer:          FindCustomerByID(alertCustomerID),
			}

			list = append(list, payment)
		}

	}

	return list
}
