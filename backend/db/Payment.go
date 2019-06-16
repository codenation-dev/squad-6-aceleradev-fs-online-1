package db

import (
	"database/sql"
	"log"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
)

//FindAllPayments retorna todos os pagamentos
func FindAllPayments(returnEmployees bool) []models.Payment {
	var (
		paymentID       int
		paymentFileName string
		paymentYear     int
		paymentMonth    int
		listEmployee    []models.Payment
	)

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
			var payment = models.Payment{
				ID:       paymentID,
				FileName: paymentFileName,
				Month:    paymentMonth,
				Year:     paymentYear}
			if returnEmployees {
				payment.EmployeePayments = findPaymentsEmployeeByPaymentID(paymentID)
			}

			listEmployee = append(listEmployee, payment)
		}

	}

	return listEmployee
}

//FindPaymentByID retorna pagamento por id
func FindPaymentByID(returnEmployees bool, ID int) models.Payment {
	var (
		paymentID       int
		paymentFileName string
		paymentYear     int
		paymentMonth    int
		payment         models.Payment
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	errQuery := db.QueryRow(
		"select pagamento.* from pagamento "+
			" where (pagame_id = $1)",
		ID).Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)

	if (errQuery != nil) && (errQuery != sql.ErrNoRows) {
		log.Println("db.FindPaymentByID()->Erro ao executar consulta. Error:", errQuery)
	}

	if paymentID > 0 {
		payment = models.Payment{
			ID:       paymentID,
			FileName: paymentFileName,
			Month:    paymentMonth,
			Year:     paymentYear}

		if returnEmployees {
			payment.EmployeePayments = findPaymentsEmployeeByPaymentID(paymentID)
		}

	}

	return payment
}

//FindPaymentByYearAndMonth retorna pagamento por ano e mes
func FindPaymentByYearAndMonth(returnEmployees bool, year int, month int) models.Payment {
	var (
		paymentID       int
		paymentFileName string
		paymentYear     int
		paymentMonth    int
		payment         models.Payment
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	errQuery := db.QueryRow(
		"select pagamento.* from pagamento "+
			" where (pagame_ano = $1) and (pagame_mes = $2)",
		year, month).Scan(&paymentID, &paymentFileName, &paymentYear, &paymentMonth)

	if (errQuery != nil) && (errQuery != sql.ErrNoRows) {
		log.Println("db.FindPaymentByYearAndMonth()->Erro ao executar consulta. Error:", errQuery)
	}

	if paymentID > 0 {
		payment = models.Payment{
			ID:       paymentID,
			FileName: paymentFileName,
			Month:    paymentMonth,
			Year:     paymentYear}
		if returnEmployees {
			payment.EmployeePayments = findPaymentsEmployeeByPaymentID(paymentID)
		}
	}

	return payment
}

//InsertPayment cadastra pagamento
func InsertPayment(returnEmployees bool, payment models.Payment) models.Payment {
	var (
		paymentID                   int
		paymentFileName             string
		paymentYear                 int
		paymentMonth                int
		paymentInserted             models.Payment
		listPaymentEmployeeInserted []models.PaymentEmployee
	)

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
			paymentEmployeeInserted := insertPaymentEmployee(db, paymentInserted, paymentEmployee)
			listPaymentEmployeeInserted = append(listPaymentEmployeeInserted, paymentEmployeeInserted)
		}
		if returnEmployees {
			paymentInserted.EmployeePayments = listPaymentEmployeeInserted
		}

	}
	return paymentInserted
}

//InsertPaymentEmployee cadastra pagamento
func insertPaymentEmployee(db *sql.DB, payment models.Payment, paymentEmployee models.PaymentEmployee) models.PaymentEmployee {
	var (
		paymentID                 int
		paymentEmployeeInserted   models.PaymentEmployee
		paymentEmployeeID         int
		paymentEmployeeName       string
		paymentEmployeeOccupation string
		paymentEmployeeDepartment string
		paymentEmployeeSalary     float64
		paymentEmployeeCustumerID int
	)

	insert :=
		`INSERT INTO public.pagamento_funcionario
		(pagame_id, pagfun_nome, pagfun_cargo, pagfun_orgao, pagfun_remuneracao, client_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		returning pagfun_id, pagame_id, pagfun_nome, pagfun_cargo, pagfun_orgao, pagfun_remuneracao, client_id;`

	var occupationFix string
	if len(paymentEmployee.Occupation) > 3 {
		occupationFix = "(vazio)"
	}
	errInsert := db.QueryRow(insert,
		payment.ID,
		paymentEmployee.Name,
		occupationFix,
		paymentEmployee.Department,
		paymentEmployee.Salary,
		paymentEmployee.Customer.ID).Scan(
		&paymentEmployeeID,
		&paymentID,
		&paymentEmployeeName,
		&paymentEmployeeOccupation,
		&paymentEmployeeDepartment,
		&paymentEmployeeSalary,
		&paymentEmployeeCustumerID)

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
			Salary:     paymentEmployeeSalary,
			Customer:   FindCustomerByID(paymentEmployeeCustumerID),
		}
	}
	return paymentEmployeeInserted
}

func findPaymentsEmployeeByPaymentID(paymentID int) []models.PaymentEmployee {
	var (
		paymentEmployeeID         int
		paymentEmployeeName       string
		paymentEmployeeOccupation string
		paymentEmployeeDepartment string
		paymentEmployeeSalary     float64
		paymentEmployeeCustumerID int

		listEmployee []models.PaymentEmployee
	)

	db := ConnectDataBase()
	defer CloseDataBase(db)

	rows, errQuery := db.Query(
		`select 
			pagamento_funcionario.pagfun_id,
			pagamento_funcionario.pagfun_nome,
			pagamento_funcionario.pagfun_cargo,
			pagamento_funcionario.pagfun_orgao,
			pagamento_funcionario.pagfun_remuneracao,
			pagamento_funcionario.client_id
		from pagamento_funcionario
		where
			pagamento_funcionario.pagame_id = $1 `, paymentID)

	if errQuery != nil {
		log.Println("db.findPaymentsEmployeeByPaymentID()->Erro ao executar consulta. Error:", errQuery)
	}

	for rows.Next() {
		err := rows.Scan(
			&paymentEmployeeID,
			&paymentEmployeeName,
			&paymentEmployeeOccupation,
			&paymentEmployeeDepartment,
			&paymentEmployeeSalary,
			&paymentEmployeeCustumerID,
		)
		if err != nil {
			log.Fatal("db.findPaymentsEmployeeByPaymentID()->Erro ao executar consulta. Error:", err)
		} else {
			var employee = models.PaymentEmployee{
				ID:         paymentEmployeeID,
				Name:       paymentEmployeeName,
				Occupation: paymentEmployeeOccupation,
				Department: paymentEmployeeDepartment,
				Salary:     paymentEmployeeSalary,
				Customer:   FindCustomerByID(paymentEmployeeCustumerID),
			}

			listEmployee = append(listEmployee, employee)
		}
	}
	return listEmployee
}

//DeletePaymentByID retona usuario pelo seu email
func DeletePaymentByID(id int) bool {
	db := ConnectDataBase()
	defer CloseDataBase(db)

	_, err1 := db.Exec(
		`update historico_alerta set pagfun_id = null 
		where pagfun_id in (select pagfun_id from pagamento_funcionario where pagame_id = $1 );`, id)

	if err1 == nil {
		_, err2 := db.Exec(`delete from pagamento_funcionario where pagame_id = $1;`, id)
		if err2 == nil {
			_, err3 := db.Exec(`delete from pagamento where pagame_id = $1;`, id)
			if err3 == nil {
				return true
			} else {
				log.Fatal("db.DeletePaymentByID()-> Error:", err3)
			}
		} else {
			log.Fatal("db.DeletePaymentByID()-> Error:", err2)
		}

	} else {
		log.Fatal("db.DeletePaymentByID()->  Error:", err1)
	}

	return false
}
