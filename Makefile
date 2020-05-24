dockermon:
	go build -mod vendor -o dockermon .

build:
	GOOS=linux GOARCH=amd64 go build -o dockermon-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o dockermon-darwin-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dockermon-linux-arm64 .

test:
	go test -mod vendor ./...

clean:
	rm dockermon
	rm dockermon-*
