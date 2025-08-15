GO = go
GOFLAGS = -trimpath
LDFLAGS = -w -s
PREFIX = /usr/local
BINDIR = $(PREFIX)/bin
TARGET = ght
SRCDIR = .
MAIN_SRC = $(SRCDIR)/main.go

.PHONY: all build clean install uninstall test fmt vet

all: build

build:
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(TARGET) $(MAIN_SRC)

clean:
	$(GO) clean
	rm -f $(TARGET)

install: build
	install -d $(BINDIR)
	install -m 755 $(TARGET) $(BINDIR)

uninstall:
	rm -f $(BINDIR)/$(TARGET)

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...
