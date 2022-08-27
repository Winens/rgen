GO = go

TARGET = rgen

.PHONY: build clean install uninstall
all: build

install:
	cp $(TARGET) /usr/bin/

uninstall:
	rm -f /usr/bin/$(TARGET)

build:
	$(GO) build -o $(TARGET)

clean:
	rm -f $(TARGET)
