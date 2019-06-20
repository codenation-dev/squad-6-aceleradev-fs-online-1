package handlers

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

var dbConsulta *sql.DB

// GetPayments retorna todos os pagamentos
func GetPayments(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if customerID, err := strconv.Atoi(c.Query("customerId")); err == nil {
		c.JSON(http.StatusOK, db.FindPayments(true, customerID))

	} else {
		c.JSON(http.StatusOK, db.FindPayments(false, 0))
	}

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
			c.AbortWithStatus(http.StatusNoContent)
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
				pathFolderExtracted := fileRarPayment[0 : len(fileRarPayment)-4]

				if _, err := os.Stat(pathFolderExtracted); !os.IsNotExist(err) {
					os.RemoveAll(pathFolderExtracted)
				}

				errExtract := services.ExtractRarFile(fileRarPayment, pathFolderExtracted)
				if errExtract == nil {
					fmt.Println("CheckPayments()-> folder extracted:" + pathFolderExtracted)

					files, err := ioutil.ReadDir(pathFolderExtracted)
					if err != nil {
						log.Fatal(err)
					}
					pathCSV := pathFolderExtracted + "/" + files[0].Name()

					fmt.Println("CheckPayments()-> TXT file:" + pathCSV)
					if _, err := os.Stat(pathCSV); err == nil {
						fmt.Println("CheckPayments()-> TXT file check: ok")

						registerPaymentsFromCSV(pathCSV, year, int(month))
					} else {
						fmt.Println("CheckPayments()-> TXT file not found! Error")
					}

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

		paymentEmployee := models.PaymentEmployee{
			ID:         0,
			Name:       line[0],
			Occupation: line[1],
			Department: line[2],
			Salary:     salary,
		}
		employeeList = append(employeeList, paymentEmployee)

		count = count + 1
	}

	dbConsulta = db.ConnectDataBase()
	beginTime := time.Now()

	jobs := make(chan models.PaymentEmployee, len(employeeList))
	results := make(chan models.PaymentEmployee, len(employeeList))

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for _, employee := range employeeList {
		jobs <- employee
	}
	close(jobs)

	var listEmployeeForRegister []models.PaymentEmployee

	for a := 1; a <= len(employeeList); a++ {
		employeeRetorno := <-results
		if (employeeRetorno.Salary >= minSalaryForRegisterPayment) || (employeeRetorno.Customer.ID > 0) {
			listEmployeeForRegister = append(listEmployeeForRegister, employeeRetorno)
			acceptPayment = acceptPayment + 1
		}
	}
	fmt.Println("registerPaymentsFromCSV()-> listEmployeeForRegister", len(listEmployeeForRegister))
	fmt.Println("registerPaymentsFromCSV()-> search customers->begin:", services.DateToStr(beginTime), "end", services.DateToStr(time.Now()))
	db.CloseDataBase(dbConsulta)
	payment.EmployeePayments = listEmployeeForRegister

	fmt.Println("registerPaymentsFromCSV()-> payment Count:", count)
	fmt.Println("registerPaymentsFromCSV()-> payment Accept:", acceptPayment)

	fmt.Println("registerPaymentsFromCSV()-> register payments in db begin")
	paymentInserted := db.InsertPayment(nil, false, payment)
	fmt.Println("registerPaymentsFromCSV()-> register payments in db end")

	if paymentInserted.ID > 0 {
		RegisterAndNotifyAlerts(paymentInserted.ID)
	}

}

func worker(id int, jobs <-chan models.PaymentEmployee, results chan<- models.PaymentEmployee) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		employee := j
		customerFound := db.FindCustomerByName(dbConsulta, j.Name)
		if customerFound.ID > 0 {
			employee.Customer = customerFound
			//fmt.Println("worker", id, "finished job", employee.Customer.ID)
		}
		//fmt.Println("worker", id, "finished job", employee.Customer.ID)
		results <- employee
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
