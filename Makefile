run:
	go run -race main.go

clean:
	go clean

build:
	@GOOS=linux GOARCH=amd64
	@go build -o gin-mongodb ./
	@echo ">> Finished"

prod:
	@./gin-mongodb
