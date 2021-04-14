run: go run main.go

build:
	@GOOS=linux GOARCH=amd64
	@go build -o gin-mongodb ./
	@echo ">> Finished"

prod:
	@./gin-mongodb
