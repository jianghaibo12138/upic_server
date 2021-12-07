build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/upic_server-linux-amd64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/upic_server-darwin-amd64
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/upic_server-windows-amd64.exe