
dev:
	nodemon --exec ENV=development go run cmd/main.go --signal SIGTERM

heroku:
	env GOOS=linux GOARCH=amd64 go build -o heroku-build/main -v cmd/main.go

push-heroku:
	/opt/homebrew/bin/git push heroku main

deploy-heroku:
	make heroku
	make push-heroku
