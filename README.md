# Scan Tools

```zsh
$ make build
```

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

![.github/workflows/main.yml](https://github.com/tz4678/scan-tools/workflows/.github/workflows/main.yml/badge.svg)

# Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять Блять
