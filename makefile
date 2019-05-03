all: linux mac
test:
	go test -v ./tests/${file}

prepare:
	go test -v ./...

## For lunux ubuntu
linux:
	make prepare
	rm -rf builds/linux
	mkdir -p builds/linux/logs
	mkdir -p builds/linux/www
	env GOOS=linux GOARCH=amd64 go build -o builds/linux/app main.go
	cp config.json builds/linux/config.json
	touch builds/linux/logs/error builds/linux/logs/info builds/linux/logs/warning
	touch builds/linux/www/index.html
	@echo "###SUCCESS BUILD LINUX ENVIRONMENT###"

## For mac os.
mac:
	make prepare
	rm -rf builds/mac
	mkdir -p builds/mac/logs
	mkdir -p builds/mac/www
	env GOOS=darwin GOARCH=amd64 go build -o builds/mac/app main.go
	cp config.json builds/mac/config.json
	touch builds/mac/logs/error builds/mac/logs/info builds/mac/logs/warning
	touch builds/mac/www/index.html
	@echo "###SUCCESS BUILD MAC ENVIRONMENT###"