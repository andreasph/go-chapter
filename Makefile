NAME=homework-$(SCOPE)
SOURCE=$(SCOPE)/main.go

all: clean build

clean:
	rm -rf build/

build:
	go build -o build/$(NAME) $(SOURCE)

fmt:
	go fmt $(shell glide novendor)

vet:
	go vet $(shell glide novendor)

lint:
	for file in $(shell find . -name '*.go' -not -path './vendor/*'); do golint $${file}; done
