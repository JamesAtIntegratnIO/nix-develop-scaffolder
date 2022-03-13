BINARY_NAME=site-map-builder

# Build for multiple architectures
build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
  