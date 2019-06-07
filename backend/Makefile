# Makefile

# postgres in docker
export DB_HOST := 172.17.0.2
export DB_USER := postgres
export DB_PASSWORD := 12345
export DB_BANCO := codenation
export DB_PORT := 5432

export JWT_SECRET = SecretJWTKeyCodeNation

run:
	go run main.go

testRoutes:	
	/usr/local/go/bin/go test -timeout 30s github.com/codenation-dev/squad-6-aceleradev-fs-online-1/backend/routes -coverprofile=/tmp/vscode-gomiv2aW/go-code-cover