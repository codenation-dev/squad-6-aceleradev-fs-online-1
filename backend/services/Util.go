package services

import (
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/models"
	"github.com/jordan-wright/email"
	"github.com/matcornic/hermes"
	"github.com/mholt/archiver"
)

//DownloadPaymentFile busca salario no site do governo
func DownloadPaymentFile(year int, month int) (string, error) {
	month = month - 1

	var months [12]string
	months[0] = "Janeiro"
	months[1] = "Fevereiro"
	months[2] = "Marco"
	months[3] = "Abril"
	months[4] = "Maio"
	months[5] = "Junho"
	months[6] = "Julho"
	months[7] = "Agosto"
	months[8] = "Setembro"
	months[9] = "Outubro"
	months[10] = "Novembro"
	months[11] = "Dezembro"

	url := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/remuneracao_" + months[month] + "_" + strconv.Itoa(year) + ".rar"

	filepath := "./temp/remuneracao_" + months[month] + "_" + strconv.Itoa(year) + ".rar"

	if _, err := os.Stat(filepath); err == nil {
		os.Remove(filepath)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		os.Remove(filepath)
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		out.Close()
		os.Remove(filepath)
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

// ExtractRarFile funcao para descompactar arquivos do tipo RAR
func ExtractRarFile(filepath string, outpath string) error {
	err := archiver.Unarchive(filepath, outpath)
	if err != nil {
		fmt.Println("Error->ExtractRarFile->", filepath, "Error: ", err)
		return err
	}

	return nil
}

//SendEmailAlertEmployeeSalary funcao que enviar notificacao para usuarios
func SendEmailAlertEmployeeSalary(listUser []models.User, listAlert []models.AlertHistory) {

	var listEmailUsers []string
	for _, user := range listUser {
		listEmailUsers = append(listEmailUsers, user.Email)
	}
	fmt.Println(listEmailUsers)

	var customersTable [][]hermes.Entry

	for _, alert := range listAlert {

		line := []hermes.Entry{
			{Key: "Cliente", Value: alert.Customer.Name},
			{Key: "Salario", Value: fmt.Sprintf("%.2f", alert.PaymentEmployee.Salary)},
		}
		customersTable = append(customersTable, line)

	}

	if (len(listEmailUsers)) > 0 {

		h := hermes.Hermes{
			Theme: new(hermes.Flat),
			Product: hermes.Product{
				Name: "Banco Uati",
				Link: "https://www.codenation.dev/acceleration/full-stack-go-react-remote-1/challenge/banco-uati",
				//Logo:      "https://www.codenation.dev/_nuxt/img/9bd98ba.svg",
				Copyright: "Copyright © 2019 CodeNation AceleraDev-Squad 6. Todos os direitos reservados.",
			},
		}

		mail := hermes.Email{
			Body: hermes.Body{
				Title:     "Olá,",
				Signature: "att",
				Intros: []string{
					"Novo pagamento processado, abaixo clientes do governo com salarios acima de R$ 20.000,00:",
				},
				Actions: []hermes.Action{
					{
						Instructions: "Clique no botao abaixo para visualizar no sistema do banco:",
						Button: hermes.Button{
							Color: "#22BC66",
							Text:  "Analisar Clientes",
							//mudar link aqui
							Link: "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
						},
					},
				},
				Table: hermes.Table{
					Data: customersTable,
					Columns: hermes.Columns{
						CustomWidth: map[string]string{
							"Cliente": "65%",
							"Salario": "35%",
						},
						CustomAlignment: map[string]string{
							"Salario": "right",
						},
					},
				},
				Outros: []string{
					"Precisa de ajuda, tem alguma duvida? Responda esse email, vamos adorar ajudar voce",
				},
			},
		}

		emailBody, err := h.GenerateHTML(mail)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}

		emailText, err := h.GeneratePlainText(mail)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}

		e := email.NewEmail()
		e.From = os.Getenv("EMAIL_SENDER_IDENTITY") + " <" + os.Getenv("EMAIL_SENDER_EMAIL") + ">"
		e.To = listEmailUsers
		e.Subject = "Alertas gerados com Pagamento do Governo SP"
		e.Text = []byte(emailText)
		e.HTML = []byte(emailBody)

		errSendMail := e.Send(os.Getenv("EMAIL_SMTP_SERVER")+":"+os.Getenv("EMAIL_SMTP_PORT"),
			smtp.PlainAuth(os.Getenv("EMAIL_SENDER_IDENTITY"),
				os.Getenv("EMAIL_SMTP_USER"),
				os.Getenv("EMAIL_SMTP_PASSWORD"),
				os.Getenv("EMAIL_SMTP_SERVER")))

		if errSendMail != nil {
			fmt.Println("Erro ao enviar email:", errSendMail)
		}

	}
}
