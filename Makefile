
dev:
	nodemon --exec ENV=development go run cmd/main.go --signal SIGTERM

heroku:
	GOARCH=amd64 go build -o heroku-build/main -v cmd/main.go
