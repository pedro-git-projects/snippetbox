run:
	go run ./cmd/web/*.go

on:
	go run ./cmd/web/*.go -addr=":$(port)"

help:
	go run ./cmd/web/*.go -help 

