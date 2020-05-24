dockermon:
	go build -mod vendor -o dockermon .

build:
	GOOS=linux GOARCH=amd64 go build -o dockermon-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o dockermon-darwin-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dockermon-linux-arm64 .

release:
	GOOS=linux GOARCH=amd64 go build -o dockermon-linux-amd64-$(shell git describe --tags --always) .
	GOOS=darwin GOARCH=amd64 go build -o dockermon-darwin-amd64-$(shell git describe --tags --always) .
	GOOS=linux GOARCH=arm64 go build -o dockermon-linux-arm64-$(shell git describe --tags --always) .


test:
	go test -mod vendor ./...

clean:
	rm dockermon*
