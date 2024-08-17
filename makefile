APP_NAME = json-converter
SRC = .
OUTPUT = $(APP_NAME).exe
LD_FLAGS = -ldflags="-H=windowsgui"

all: build

build:
	go build $(LD_FLAGS) -o $(OUTPUT) $(SRC)

clean:
	rm -f $(OUTPUT)

run: build
	./$(OUTPUT)

.PHONY: all build clean run
