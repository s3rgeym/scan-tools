# Scan Tools

```zsh
$ go run cmd/<binary>
$ go build cmd/<binary>
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
