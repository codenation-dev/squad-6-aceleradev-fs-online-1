# Gestão de clientes Banco Uati

## Instruções

 - entrar na pasta "go/src", no meu caso "/home/rui/go/src/"

   `git clone https://github.com/ruiblaese/projeto-codenation-banco-uati`   
   `cd projeto-codenation-banco-uati`   

   
- baixar dependencias
   `go get github.com/gin-gonic/gin`   
   `go get github.com/gin-gonic/contrib/static`   
   `go get github.com/appleboy/gin-jwt`
   `go get github.com/lib/pq`

 - criar sua versao do Makefile.exemple
   `cp Makefile.example Makefile`
 - criar ou iniciar docker com postgresql
 - criar banco de dados no postgres: codenation
 - executar script com estrutura basica do banco (.\\assets\\estrutura.sql)
 - executar projeto
`make run`

 - acessar: http://localhost:3000


## Docker
####Subindo com Docker
`$ docker pull postgres`   
`$ docker volume create pgdata`   
`$ docker run --name postgres -e POSTGRES_PASSWORD=12345 -v`   
`pgdata:/var/lib/postgresql/data -d postgres`

#####Descobrindo Ip do Docker
Para gerenciar nosso postgres podemos usar o psql:
Vamos descobrir o IP de nosso server postgres
`$ docker inspect postgres | grep IPAddress`

retorno parecido com esse:
`"SecondaryIPAddresses": null,`
`"IPAddress": "172.17.0.2",`

###Iniciando docker pela segunda vez
`docker start postgres`

## Recomendações
  - Microsoft Visual Studio Code -> https://code.visualstudio.com/download
  - insomnia -> https://insomnia.rest/download/
  - docker