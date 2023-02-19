BINARY_NAME=protovision

test:
	go test -v main.go

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux_amd64 main.go
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}-linux_arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin_amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows_amd64 main.go
