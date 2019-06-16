package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/services"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/db"
	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
	"github.com/gin-gonic/gin"

	"golang.org/x/text/encoding/charmap"
)

// GetPayments retorna todos os pagamentos
func GetPayments(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, db.FindAllPayments(true))
}

// GetPayment retornoa um pagamento
func GetPayment(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var payment models.Payment

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		payment = db.FindPaymentByID(true, id)

		if payment.ID > 0 {
			c.JSON(http.StatusOK, payment)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}

	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// DeletePayment deleta pagamento
func DeletePayment(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		apagou := db.DeletePaymentByID(id)

		if apagou {
			c.JSON(http.StatusOK, gin.H{"code": "OK", "message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "Invalid param"})
	}
}

// GetCheckPayments inicia busca e processamento de novos pagamentos
func GetCheckPayments(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	go CheckPayments()
	c.JSON(http.StatusProcessing, gin.H{
		"message": "start checking"})

}

// CheckPayments verifica se existe pagamentos para baixar e processar
func CheckPayments() {

	currentTime := time.Now()

	run := true
	for run {
		run = false
		year := currentTime.Year()
		month := currentTime.Month()

		payment := db.FindPaymentByYearAndMonth(false, year, int(month))

		if payment.ID <= 0 {
			fileRarPayment, errDownload := services.DownloadPaymentFile(year, int(month))
			if errDownload == nil {
				fmt.Println("CheckPayments()-> file downloaded", fileRarPayment)
				pathFolderCSV := fileRarPayment[0 : len(fileRarPayment)-4]

				if _, err := os.Stat(pathFolderCSV); !os.IsNotExist(err) {
					os.RemoveAll(pathFolderCSV)
				}

				errExtract := services.ExtractRarFile(fileRarPayment, pathFolderCSV)
				if errExtract == nil {
					fmt.Println("CheckPayments()-> folder extracted:" + pathFolderCSV)
					pathCSV :=
						pathFolderCSV +
							strings.Replace(
								fileRarPayment[6:len(fileRarPayment)-4], "remuneracao", "Remuneracao", -1) +
							".txt"

					fmt.Println("CheckPayments()-> CSV file:" + pathCSV)
					if _, err := os.Stat(pathCSV); err == nil {
						fmt.Println("CheckPayments()-> CSV file check: ok")
					}
					registerPaymentsFromCSV(pathCSV, year, int(month))
				} else {
					fmt.Println("CheckPayments()-> error to extract ->", fileRarPayment)
				}
			} else {
				fmt.Println("CheckPayments()-> error on download ->", year, "-", int(month))
				//processa novamente procurando no mes anterior a ultima busca
				run = true
				currentTime = currentTime.AddDate(0, -1, 0)
			}
		} else {
			fmt.Println("CheckPayments()-> not found new payment")
			fmt.Println("CheckPayments()-> last payment register ->", year, "-", int(month))
		}
	}
}

func iso88591toUtf8(fileNameIso88591 string, fileNameOutUtf8 string) {

	fmt.Println("iso88591toUtf8()-> file decode:", fileNameIso88591, "to", fileNameOutUtf8)

	f, err := os.Open(fileNameIso88591)
	if err != nil {
		// handle file open error
	}
	out, err := os.Create(fileNameOutUtf8)
	if err != nil {
		// handler error
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)
	io.Copy(out, r)
	out.Close()
	f.Close()
}

func registerPaymentsFromCSV(fileName string, year int, month int) {

	payment := models.Payment{
		FileName: fileName,
		Year:     year,
		Month:    month}

	fileNameCsvUtf8 := strings.Replace(fileName, ".txt", ".csv", -1)
	iso88591toUtf8(fileName, fileNameCsvUtf8)

	f, err := os.Open(fileNameCsvUtf8)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()

	minSalaryForRegisterPayment, err := strconv.ParseFloat(os.Getenv("CONFIG_MIN_SALARY_REGISTER_PAYMENT"), 64)
	if err != nil {
		minSalaryForRegisterPayment = 20000
	}

	count := 0
	acceptPayment := 0

	var employeeList []models.PaymentEmployee
	for _, line := range lines {
		var salary float64
		if salary, err = strconv.ParseFloat(strings.ReplaceAll(line[3], ",", "."), 64); err != nil {
			salary = 0.0
		}

		//busca para ver se eh clietne do banco
		//customer := db.FindCustomerByName(line[0])

		//busca desativada porque eh muito lento, se sobrar tempo melhorar isso
		customer := models.Customer{}

		if (salary >= minSalaryForRegisterPayment) || (customer.ID > 0) {

			//busca ativada apenas para quem ja ganha acima de x valor
			customer := db.FindCustomerByName(line[0])
			if customer.ID > 0 {
				paymentEmployee := models.PaymentEmployee{
					ID:         0,
					Name:       line[0],
					Occupation: line[1],
					Department: line[2],
					Salary:     salary,
					Customer:   customer,
				}
				employeeList = append(employeeList, paymentEmployee)
				acceptPayment = acceptPayment + 1
			}
		}
		count = count + 1
	}
	payment.EmployeePayments = employeeList

	fmt.Println("registerPaymentsFromCSV()-> payment Count:", count)
	fmt.Println("registerPaymentsFromCSV()-> payment Accept:", acceptPayment)

	fmt.Println("registerPaymentsFromCSV()-> register payments in db begin")
	paymentInserted := db.InsertPayment(false, payment)
	fmt.Println("registerPaymentsFromCSV()-> register payments in db end")

	if paymentInserted.ID > 0 {
		RegisterAndNotifyAlerts(paymentInserted.ID)
	}
}

//MonitorPayments monitora pagamentos a cada 10 horas
func MonitorPayments() {

	go CheckPayments()

	nextTime := time.Now()
	//teste 2 seconds
	//nextTime = nextTime.Add(time.Second * 2)

	nextTime = nextTime.AddDate(0, 0, 1)
	fmt.Println("MonitorPayments()-> next search", nextTime)
	time.Sleep(time.Until(nextTime))

	go MonitorPayments()
}
