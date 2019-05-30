# Makefile

# postgres in docker
export DB_HOST := localhost
export DB_USER := postgres
export DB_PASSWORD := cl123456
export DB_BANCO := codenation
export DB_PORT := 5432

export JWT_SECRET = SecretJWTKeyCodeNation

run:
	go run main.go