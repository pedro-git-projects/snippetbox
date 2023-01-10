run:
	go run ./cmd/web/*.go -dsn="web:$(pass)@/snippetbox?parseTime=true"

on:
	go run ./cmd/web/*.go -addr=":$(port)"

help:
	go run ./cmd/web/*.go -help 

build:
	go build -o build/ ./cmd/web/

bin:
	./build/web -dsn="web:$(pass)@/snippetbox?parseTime=true"
