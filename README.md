# ght
go-http-get  

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
sudo mv ght /usr/local/bin
sudo chown root:root /usr/local/bin/ght
which ght
```

## Usage
URL in `$1`.
```bash
$ ght
Usage: ght "https://google.com/"
```

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
