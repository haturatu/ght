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

or you can use the Makefile to build it.

```bash
make
```

## Install
You can install it to `/usr/local/bin` with the following command.

```bash
sudo make install
```
## Uninstall

```bash
sudo make uninstall
```

## Usage

URL in `$1`.

```bash
$ ght --help
usage: ght [-h|--help] [-u|--url "<value>"] [-m|--markdown] [-c|--copy]

           Get HTML Title

Arguments:

  -h  --help      Print help information
  -u  --url       URL to fetch
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

## Makefile

-   `make build`: Build the binary.
-   `make install`: Install the binary to `/usr/local/bin`.
-   `make uninstall`: Uninstall the binary.
-   `make clean`: Remove the binary.
-   `make fmt`: Format the code.
-   `make vet`: Run go vet.
-   `make test`: Run tests.

