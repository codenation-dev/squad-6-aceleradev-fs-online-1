package services

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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

//SendEmailNotificationNewsEmployee funcao que enviar notificacao para usuarios
func SendEmailNotificationNewsEmployee() {

	h := hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name:      "Banco Uati",
			Link:      "https://www.codenation.dev/acceleration/full-stack-go-react-remote-1/challenge/banco-uati",
			Logo:      "https://www.codenation.dev/_nuxt/img/9bd98ba.svg",
			Copyright: "Copyright © 2019 CodeNation AceleraDev-Squad 6. Todos os direitos reservados.",
		},
	}

	mail := hermes.Email{
		Body: hermes.Body{
			Title:     "Olá,",
			Signature: "att",
			Intros: []string{
				"Novo pagamento processado, abaixo clientes do governo com salarios acima de 20 mil:",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Clique no botao abaixo para visualizar no sistema do banco:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					// List of rows
					{
						{Key: "Cliente", Value: "Open source programming language "},
						{Key: "Salario", Value: "$10.99"},
					},
					{

						{Key: "Cliente", Value: "Programmatically "},
						{Key: "Salario", Value: "$1.99"},
					},
				},
				Columns: hermes.Columns{
					// Custom style for each rows
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

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(mail)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(mail)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = ioutil.WriteFile("./temp/preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = ioutil.WriteFile("./temp/preview.text", []byte(emailText), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
	/*
		e := email.NewEmail()
		e.From = os.Getenv("EMAIL_SENDER_IDENTITY") + " <" + os.Getenv("EMAIL_SENDER_EMAIL") + ">"
		e.To = []string{"ruiblaese@gmail.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte(emailText)
		e.HTML = []byte(emailBody)
		errE := e.Send(os.Getenv("EMAIL_SMTP_SERVER")+":"+os.Getenv("EMAIL_SMTP_PORT"),
			smtp.PlainAuth(os.Getenv("EMAIL_SENDER_IDENTITY"),
				os.Getenv("EMAIL_SMTP_USER"),
				os.Getenv("EMAIL_SMTP_PASSWORD"),
				os.Getenv("EMAIL_SMTP_SERVER")))

		fmt.Println(errE)
		fmt.Println("-")

		fmt.Println(os.Getenv("EMAIL_SMTP_SERVER"))
		fmt.Println(os.Getenv("EMAIL_SMTP_PORT"))
		fmt.Println(os.Getenv("EMAIL_SMTP_USER"))
		fmt.Println(os.Getenv("EMAIL_SMTP_PASSWORD"))
		fmt.Println(os.Getenv("EMAIL_SENDER_EMAIL"))
		fmt.Println(os.Getenv("EMAIL_SENDER_IDENTITY"))
	*/
}
