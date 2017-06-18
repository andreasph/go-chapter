NAME=homework-$(SCOPE)
SOURCE=$(SCOPE)/main.go
DEPEND=github.com/Masterminds/glide

all: clean depend build

clean:
	rm -rf build/

build:
	go build -o build/$(NAME) $(SOURCE)

depend:
	go get -u -v $(DEPEND)
	cd $(SCOPE) && glide install

fmt:
	go fmt $(shell glide novendor)

vet:
	go vet $(shell glide novendor)

lint:
	for file in $(shell find . -name '*.go' -not -path './vendor/*'); do golint $${file}; done

all: clean depend build