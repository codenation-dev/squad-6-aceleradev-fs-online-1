package db

import (
	"log"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
)

//InsertAlertHistory cadastra historico de alerta e retorna registro inserido no banco
func InsertAlertHistory(alertHistory models.AlertHistory) models.AlertHistory {
	var (
		ID                   int
		alertHistoryInserted models.AlertHistory
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	insert :=
		`INSERT INTO public.historico_alerta
		(hisale_data, usuari_id, client_id, pagfun_id)
		VALUES ($1, $2, $3, $4) 
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
