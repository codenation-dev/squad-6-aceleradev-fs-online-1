package db

import (
	"database/sql"
	"log"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/models"
)

var (
	paymentID       int
	paymentFileName string
	paymentYear     int
	paymentMonth    int
)
var (
	paymentEmployeeID         int
	paymentEmployeeName       string
	paymentEmployeeOccupation string
	paymentEmployeeDepartment string
	paymentEmployeeSalary     float64
)

//FindAllPayments retorna todos os pagamentos
func FindAllPayments(returnEmployees bool) []models.Payment {
	var listUsers []models.Payment

	db := ConnectDataBase()
	defer CloseDataBase(db)

	rows, errQuery := db.Query("select pagamento.* from pagamento")
	if errQuery != nil {
		log.Println("db.FindAllPayments()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {
		err := rows.Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)
		if err != nil {
			log.Fatal("db.FindAllPayments()->Erro ao executar consulta. Error:", err)
		} else {
			var user = models.Payment{
				ID:       userID,
				FileName: paymentFileName,
				Month:    paymentMonth,
				Year:     paymentYear}

			listUsers = append(listUsers, user)
		}

	}

	return listUsers
}

//FindPaymentByID retorna pagamento por id
func FindPaymentByID(returnEmployees bool, ID int) models.Payment {
	var payment models.Payment

	db := ConnectDataBase()
	defer CloseDataBase(db)

	errQuery := db.QueryRow(
		"select pagamento.* from pagamento "+
			" where (pagame_id = $1)",
		ID).Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)

	if paymentID > 0 {
		payment = models.Payment{
			ID:       userID,
			FileName: paymentFileName,
			Month:    paymentMonth,
			Year:     paymentYear}
	}

	if (errQuery != nil) && (errQuery != sql.ErrNoRows) {
		log.Println("db.FindPaymentByYearAndMonth()->Erro ao executar consulta. Error:", errQuery)
	}

	return payment
}

//FindPaymentByYearAndMonth retorna pagamento por ano e mes
func FindPaymentByYearAndMonth(returnEmployees bool, year int, month int) models.Payment {
	var payment models.Payment

	db := ConnectDataBase()
	defer CloseDataBase(db)

	errQuery := db.QueryRow(
		"select pagamento.* from pagamento "+
			" where (pagame_ano = $1) and (pagame_mes = $2)",
		year, month).Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)

	if paymentID > 0 {
		payment = models.Payment{
			ID:       paymentID,
			FileName: paymentFileName,
			Month:    paymentMonth,
			Year:     paymentYear}
	}

	if (errQuery != nil) && (errQuery != sql.ErrNoRows) {
		log.Println("db.FindPaymentByYearAndMonth()->Erro ao executar consulta. Error:", errQuery)
	}

	return payment
}

//InsertPayment cadastra pagamento
func InsertPayment(returnEmployees bool, payment models.Payment) models.Payment {
	var paymentInserted models.Payment

	db := ConnectDataBase()
	defer CloseDataBase(db)

	insert :=
		`INSERT INTO public.pagamento
		(pagame_arquivo, pagame_ano, pagame_mes)
		VALUES ($1, $2, $3) returning pagame_id, pagame_arquivo, pagame_ano, pagame_mes;`

	errInsert := db.QueryRow(insert,
		payment.FileName, payment.Year, payment.Month).Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)

	if (errInsert != nil) && (errInsert != sql.ErrNoRows) {
		log.Println("db.InsertPayment->Erro ao executar insert. Error:", errInsert)
	} else {
		paymentInserted = models.Payment{
			ID:       paymentID,
			FileName: paymentFileName,
			Year:     paymentYear,
			Month:    paymentMonth}

		for _, paymentEmployee := range payment.EmployeePayments {
			insertPaymentEmployee(db, paymentInserted, paymentEmployee)
		}

	}

	return paymentInserted
}

//InsertPaymentEmployee cadastra pagamento
func insertPaymentEmployee(db *sql.DB, payment models.Payment, paymentEmployee models.PaymentEmployee) models.PaymentEmployee {
	var paymentEmployeeInserted models.PaymentEmployee

	insert :=
		`INSERT INTO public.pagamento_funcionario
		(pagame_id, pagfun_nome, pagfun_cargo, pagfun_orgao, pagfun_remuneracao)
		VALUES ($1, $2, $3, $4, $5)
		returning pagfun_id, pagame_id, pagfun_nome, pagfun_cargo, pagfun_orgao, pagfun_remuneracao;`

	var occupationFix string
	if len(paymentEmployee.Occupation) > 3 {
		occupationFix = "(vazio)"
	}
	errInsert := db.QueryRow(insert,
		payment.ID, paymentEmployee.Name, occupationFix, paymentEmployee.Department,
		paymentEmployee.Salary).Scan(&paymentEmployeeID, &paymentID, &paymentEmployeeName,
		&paymentEmployeeOccupation, &paymentEmployeeDepartment, &paymentEmployeeSalary)

	if (errInsert != nil) && (errInsert != sql.ErrNoRows) {
		log.Println("db.InsertPaymentEmployee->Erro ao executar insert. Error:",
			errInsert, paymentEmployee,
			"\n", "("+paymentEmployee.Name+","+paymentEmployee.Department+","+paymentEmployee.Occupation+")")
	} else {
		paymentEmployeeInserted = models.PaymentEmployee{
			ID:         paymentEmployeeID,
			Name:       paymentEmployeeName,
			Occupation: paymentEmployeeOccupation,
			Department: paymentEmployeeDepartment,
			Salary:     paymentEmployeeSalary}
	}
	return paymentEmployeeInserted
}
