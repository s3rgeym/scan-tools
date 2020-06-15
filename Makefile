export GO111MODULE=on

build:
	for item in cmd/*; do \
		go build -v -o "build/$${item##*/}" "$$item/main.go"; \
		# можно сборки под разные архитектуры сделать: \
		# GOOS=linux GOARCH=amd64 go build -v -o "build/$${item##*/}-linux-amd64" "$$item/main.go"; \
		# GOOS=windows GOARCH=amd64 go build -v -o "build/$${item##*/}-win-amd64.exe" "$$item/main.go"; \
		# GOOS=darwin GOARCH=amd64 go build -v -o "build/$${item##*/}-darwin-amd64" "$$item/main.go"; \
	done

clean:
	rm -rf build/*
