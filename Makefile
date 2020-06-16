PREFIX ?= /usr/local
# без этого флага не компилируется
export GO111MODULE = on

build:
	for item in cmd/*; do \
		go build -v -o "bin/$${item##*/}" "$$item/main.go"; \
		# можно сборки под разные архитектуры сделать: \
		# GOOS=linux GOARCH=amd64 go build -v -o "bin/$${item##*/}-linux-amd64" "$$item/main.go"; \
		# GOOS=windows GOARCH=amd64 go build -v -o "bin/$${item##*/}-win-amd64.exe" "$$item/main.go"; \
		# GOOS=darwin GOARCH=amd64 go build -v -o "bin/$${item##*/}-darwin-amd64" "$$item/main.go"; \
	done

clean:
	rm -rf bin/*

install:
	install -m 755 ./bin/* $(PREFIX)/bin

uninstall:
	for item in cmd/*; do \
		rm $(PREFIX)/bin/$${item##*/}; \
	done \

# если бы в каталоге был файл install, то `make install` запустило бы его без
# этой настройки
.PHONY: build clean install uninstall
