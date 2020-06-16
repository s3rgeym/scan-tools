# Scan Tools

![Release workflow](https://github.com/tz4678/scan-tools/workflows/Release%20workflow/badge.svg)

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
$ cat ./top-1m.csv | head -n 100 | cut -f2 -d, | addscheme
https://google.com
https://youtube.com
https://facebook.com
...
# Usage:
$ addsheme -h
```
