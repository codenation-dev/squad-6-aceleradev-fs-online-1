# Gestão de clientes Banco Uati

## Instruções

 - entrar na pasta "go/src", no meu caso "/home/rui/go/src/"   
   `git clone https://github.com/codenation-dev/squad-6-aceleradev-fs-online-1`   
   `cd squad-6-aceleradev-fs-online-1`   
   
- baixar dependencias   
   `go get github.com/gin-gonic/gin`   
   `go get github.com/gin-gonic/contrib/static`   
   `go get github.com/appleboy/gin-jwt`
   `go get github.com/lib/pq`

 - criar sua versao do Makefile.exemple    
   `cp Makefile.example Makefile`

 - criar\iniciar docker(postgresql) ou proprio postgresql
 - criar banco de dados no postgres: codenation
 - executar script com estrutura basica do banco (.\\assets\\estrutura.sql)

 - executar projeto
`make run`

 - acessar: http://localhost:3000

## Rotas
 - Login: http://localhost:3000/api/v1/signin
 `enviar header com token nas proximas rotas`
##### Usuario
 - GET: http://localhost:3000/api/v1/user/
 - GET (Id ou Email): http://localhost:3000/api/v1/user/:idOuEmail
 - PUT: http://localhost:3000/api/v1/user/:id
 - POST: http://localhost:3000/api/v1/user/  
 - DELETE: http://localhost:3000/api/v1/user/:id   

##### Cliente
 - GET: http://localhost:3000/api/v1/customer/
 - GET (Id ou Email): http://localhost:3000/api/v1/customer/:id
 - PUT: http://localhost:3000/api/v1/customer/:id
 - POST: http://localhost:3000/api/v1/customer/  
 - DELETE: http://localhost:3000/api/v1/customer/:id   

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

## Recomendações
  - [Microsoft Visual Studio Code](https://code.visualstudio.com/download)
  - [Insomnia](https://insomnia.rest/download/)
  - [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## Links
  - [Projeto](https://github.com/codenation-dev/squad-6-aceleradev-fs-online-1)
  - [Chat CodeNation](https://chat.codenation.com.br/)
  - [Chat CodeNation Nosso Squad](https://chat.codenation.com.br/group/fullstack-remote-1-squad6-v2)
  - [Documentacao de configuracao de ambiente da CodeNation](https://drive.google.com/file/d/1639-YzRhVUEHHbh5E-2u6mjDHeNPKzP1/view)
  
  - [02 de maio | 1º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=3iUf0jk2IzY)

- [08 de maio | 2º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=022c5nk-5oI)

- [15 de maio | 3º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=1Aa7mabvWRk)

- [22 de maio | 4º Workshop - AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=zZGT_DyYYSo)

- [23 de maio | Webservice REST em Go com TDD | AceleraDev Full Stack Go + React - Codenation](https://www.youtube.com/watch?v=5I_IqL4Jb9U)
