.PHONY: clean
clean:
	rm please-linux-amd64 ||:
	rm please-windows-amd64.exe ||:
	rm please-linux-arm ||:

please-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o please-linux-amd64 main.go
please-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o please-windows-amd64.exe main.go
please-linux-arm:
	GOOS=linux GOARCH=arm go build -o please-linux-arm main.go

.PHONY: install
install: please-linux-amd64
	mv please-linux-amd64 ~/.local/bin/please


.PHONY: all
all: clean please-linux-amd64 please-linux-arm please-windows-amd64