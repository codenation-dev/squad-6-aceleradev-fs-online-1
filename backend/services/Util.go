package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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
