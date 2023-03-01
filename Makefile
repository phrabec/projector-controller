BINARY_NAME=projector-controller

build:
	GOARCH=arm GOARM=7 GOOS=linux go build -o bin/${BINARY_NAME}-rpi main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}.exe main.go
	GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-macos main.go

clean:
	go clean
	rm bin/*

vet:
	go vet
