# Scan Tools

![.github/workflows/release.yml](https://github.com/tz4678/scan-tools/workflows/.github/workflows/release.yml/badge.svg)

## Download Binaries

Download `release.zip` from [latest release](https://github.com/tz4678/scan-tools/releases/latest).

## Installation

```zsh
$ make clean
$ make build
$ sudo make install
$ sudo make uninstall
```

## Tests

```zsh
$ go test ./cmd/* -v
$ go test ./... -v
$ go test ./cmd/* ./... -v
```

### Tools

## addscheme

Add scheme to host:

```zsh
$ wget http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
$ unzip top-1m.csv.zip
$ rm -rf top-1m.csv.zip
$ cat ./top-1m.csv | head -n 100 | cut -f2 -d, | go run ./cmd/addscheme
https://google.com
https://youtube.com
https://facebook.com
...
```
