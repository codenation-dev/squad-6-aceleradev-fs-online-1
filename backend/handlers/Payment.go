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
				fmt.Println("arquivo compactado:" + fileRarPayment)
				pathFolderCSV := fileRarPayment[0 : len(fileRarPayment)-4]

				if _, err := os.Stat(pathFolderCSV); !os.IsNotExist(err) {
					os.RemoveAll(pathFolderCSV)
				}

				errExtract := services.ExtractRarFile(fileRarPayment, pathFolderCSV)
				if errExtract == nil {
					fmt.Println("pasta extraida:" + pathFolderCSV)
					pathCSV :=
						pathFolderCSV +
							strings.Replace(
								fileRarPayment[6:len(fileRarPayment)-4], "remuneracao", "Remuneracao", -1) +
							".txt"

					fmt.Println("arquivo csv:" + pathCSV)
					if _, err := os.Stat(pathCSV); err == nil {
						fmt.Println("arquivo csv ok")
					}
					registerPaymentsFromCSV(pathCSV, year, int(month))
				} else {
					fmt.Println("CheckPayments()-> erro ao descompactar ->", fileRarPayment)
				}
			} else {
				fmt.Println("CheckPayments()-> erro ao fazer download ->", year, "-", int(month))
				//processa novamente procurando no mes anterior a ultima busca
				run = true
				currentTime = currentTime.AddDate(0, -1, 0)
			}
		}
	}
}

func iso88591toUtf8(fileNameIso88591 string, fileNameOutUtf8 string) {
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

	// Open CSV file
	f, err := os.Open(fileNameCsvUtf8)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()

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
	}
	payment.EmployeePayments = employeeList

	fmt.Println("inicia cadastro pagamento")
	db.InsertPayment(false, payment)
	fmt.Println("fim")
}
