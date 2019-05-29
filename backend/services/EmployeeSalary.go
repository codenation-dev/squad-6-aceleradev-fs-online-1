package services

import "net/http"

//GetEmployeeSalary busca salario no site do governo
func GetEmployeeSalary(nameEmployee string) {

	var res, _ = http.Get("http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx?ScriptManager1=ScriptManager1%7CbtnExibirRelatorio&__EVENTTARGET=&__EVENTARGUMENT=&__LASTFOCUS=&__VIEWSTATE=%aaaa&txtNome=ANTONIO%20MARCOS%20SIQUEIRA&orgao=-1&cargo=-1&situacao=-1&txtDe=&txtAte=&hdInicio=&hdFinal=&hdPaginaAtual=&hdTotal=&__ASYNCPOST=true&btnExibirRelatorio=Pesquisar")

	println(res)

	println(nameEmployee)
}
