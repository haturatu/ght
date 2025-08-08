# ght
get-http-title  

## Build
Please check if the Go path is running.
```bash
which go
```

requirements:
```bash
go get github.com/atotto/clipboard
go get github.com/akamensky/argparse
```
  
And then compile.
```bash
go build -o ght main.go
```
Move the executable binary to `/usr/local/bin` to work with CLI.
```bash
sudo mv ght /usr/local/bin/
sudo chown root:root /usr/local/bin/ght
which ght
```

## Usage
URL in `$1`.
```bash
$ ght
[--url] is required
usage: ght [-h|--help] --url "<value>" [-m|--markdown] [-c|--copy]

           Get HTML Title

Arguments:

  -h  --help      Print help information
      --url       URL to fetch
  -m  --markdown  Output in Markdown format
  -c  --copy      Copy to clipboard

```
exec
```bash
$ ght "https://google.com/"
Google
```
Copy to clipboard
```bash
$ ght -mc "https://google.com/"
[Google](https://google.com/)
```

Done!

### for example
If you have a file called `urls` with URLs listed.
```bash:urls
https://www.google.com/
https://soulminingrig.com/
https://soulminingrig.com/ab/
```
Single proccessing
```bash
cat urls | while read -r url ; do ght $url ; done
```
Parallel processing with `xargs`
```bash
cat urls | xargs -P 4 -I {} ght {}
```
