# ght
get-http-title  

## Build
Please check if the Go path is running.
```bash
which go
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
usage: ght [<flags>] <url>

Get HTML Title


Flags:
  -h, --[no-]help      Show context-sensitive help (also try --help-long and --help-man).
  -m, --[no-]markdown  Output in Markdown format
  -c, --[no-]copy      Copy to clipboard

Args:
  <url>  URL to fetch
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
