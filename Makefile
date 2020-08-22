NAME=mvc
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
export GO111MODULE=on

.PHONY: vet test deps cover

build: test
	go build -ldflags "-X main.gitcommit=$(CURRENT_REVISION)" -o build/$(NAME) ./cmd/$(NAME)

run: build
	./build/$(NAME) -d

test: vet
	go test -v -short ./...

vet: deps
	go vet -all .

deps: 
	go get -d -v -t ./...

testdeps:
	go get golang.org/x/tools/cmd/cover
	go get github.com/pierrre/gotestcover
	go get github.com/mattn/goveralls

cover: testdeps
	gotestcover -v -covermode=count -coverprofile=.profile.cov -parallelpackages=4 ./...
