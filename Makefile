run:
	go run main.go
test:
	go test -v ./... -cover

gen:
	go generate ./...