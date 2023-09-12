.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

# для запуска: make run
run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot:v0.1 

start-container:
	docker run --name telegram-bot -p 80:80 --env-file .env telegram-bot:v0.1 