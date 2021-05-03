build:
	go build -o ./cmd/server/main.go
compile:
	GOOS=linux GOARCH=amd64 go build -v github.com/1r0npipe/search-In-CSV
	GOOS=windows GOARCH=amd64 go build -v github.com/1r0npipe/search-In-CSV
	GOOS=darwin GOARCH=amd64 go build -v github.com/1r0npipe/search-In-CSV
run:
	go run ./cmd/server/main.go
