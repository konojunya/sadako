build:
	go build -o speak cmd/speak.go

install:
	go build -o sadako main.go
	chmod u+x sadako