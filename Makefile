test:
	go build -v ./...
run:
	swag init && go run .