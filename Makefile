BINARY_NAME=projector-controller

build:
	GOARCH=arm GOARM=7 GOOS=linux go build -o ${BINARY_NAME}-rpi main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}.exe main.go

clean:
	go clean
	rm ${BINARY_NAME}-rpi
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

vet:
	go vet
