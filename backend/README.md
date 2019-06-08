# Gestão de clientes Banco Uati

## Instruções

 - entrar na pasta "go/src/github/", no meu caso "/home/rui/go/src/github/"   
   `git clone https://github.com/codenation-dev/squad-6-aceleradev-fs-online-1/`   
   `cd squad-6-aceleradev-fs-online-1`   
   `cd backend`   
   
- baixar dependencias   
   `go get`   
   `go get ./...`       

 - criar sua versao do Makefile.exemple    
   `cp .env.example .env`

 - criar\iniciar docker(postgresql) ou proprio postgresql
 - criar banco de dados no postgres: codenation
 - executar script com estrutura basica do banco (.\\db\\migrations\\)

 - executar projeto
`make run`

 - acessar: http://localhost:4000

## Rotas
 - Login: http://localhost:4000/api/v1/signin `retorno do login deve ser usado header\token nas proximas rotas`
##### Usuario
 - GET: http://localhost:4000/api/v1/user/
 - GET (Id ou Email): http://localhost:4000/api/v1/user/:idOuEmail
 - PUT: http://localhost:4000/api/v1/user/:id
 - POST: http://localhost:4000/api/v1/user/  
 - DELETE: http://localhost:4000/api/v1/user/:id   

##### Cliente
 - GET: http://localhost:4000/api/v1/customer/
 - GET (Id ou Email): http://localhost:4000/api/v1/customer/:id
 - PUT: http://localhost:4000/api/v1/customer/:id
 - POST: http://localhost:4000/api/v1/customer/  
 - DELETE: http://localhost:4000/api/v1/customer/:id   

 - POST: http://localhost:4000/api/v1/customer/upload `formulario de testes: http://localhost:4000/upload.html`   

## Docker
#### Subindo com Docker
`$ docker pull postgres`   
`$ docker volume create pgdata`   
`$ docker run --name postgres -e POSTGRES_PASSWORD=12345 -v pgdata:/var/lib/postgresql/data -d postgres`

##### Descobrindo Ip do Docker
Para gerenciar nosso postgres podemos usar o psql:
Vamos descobrir o IP de nosso server postgres
`$ docker inspect postgres | grep IPAddress`

retorno parecido com esse:
`"SecondaryIPAddresses": null,`
`"IPAddress": "172.17.0.2",`

### Iniciando docker pela segunda vez
`docker start postgres`

## Links
  - [Projeto](https://github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend)
  - [Chat CodeNation](https://chat.codenation.com.br/)
  - [Chat CodeNation Nosso Squad](https://chat.codenation.com.br/group/fullstack-remote-1-squad6-v2)
  - [Documentacao de configuracao de ambiente da CodeNation](https://drive.google.com/file/d/1639-YzRhVUEHHbh5E-2u6mjDHeNPKzP1/view)
  
  - [02 de maio | 1º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=3iUf0jk2IzY)

- [08 de maio | 2º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=022c5nk-5oI)

- [15 de maio | 3º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=1Aa7mabvWRk)

- [22 de maio | 4º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=zZGT_DyYYSo)

- [23 de maio | Webservice REST em Go com TDD | AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=5I_IqL4Jb9U)
