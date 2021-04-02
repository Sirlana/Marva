all: linux mac
test:
	go run -v ./tests/${file}

prepare:
	go test -v ./...

## For lunux ubuntu
linux:
	make prepare
	rm -rf build/linux
	mkdir -p build/linux/logs
	mkdir -p build/linux/www
	env GOOS=linux GOARCH=amd64 go build -o build/linux/app main.go
	cp config.sir build/linux/config.sir
	touch build/linux/logs/error build/linux/logs/info build/linux/logs/warning
	touch build/linux/www/index.html

## For mac os.
mac:
	make prepare
	rm -rf build/mac
	mkdir -p build/mac/logs
	mkdir -p build/linux/www
	env GOOS=darwin GOARCH=amd64 go build -o build/mac/app main.go
	cp config.sir build/mac/config.sir
	touch build/mac/logs/error build/mac/logs/info build/mac/logs/warning
	touch build/linux/www/index.html