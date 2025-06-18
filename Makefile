generate:
	go generate ./ent

docs:
	swag init

run:
	go run main.go

mocks:
	mockery

test:
	go test -v ./...