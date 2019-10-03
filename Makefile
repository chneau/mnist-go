.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY:

name=$(shell basename $(CURDIR))
commit_hash=$(shell git log -1 --pretty=format:"%H")
commit_date=$(shell git log -1 --pretty=format:"%cD")
commit_timestamp=$(shell git log -1 --pretty=format:"%at")
current_dir = $(shell pwd)
target_dir = /script
port = 8080
ip = \*

run: build exec clean

exec:
	./bin/${name}

build:
	CGO_ENABLED=0 go build -trimpath -o bin/${name} -ldflags '-s -w -extldflags "-static"'

clean:
	rm -rf bin

test:
	go test -cover -count=1 ./...

deps:
	govendor init
	govendor add +e
	govendor update +v

dev:
	go get -u -v github.com/kardianos/govendor

spark:
	docker run --rm -it --name=kaggle -p=$(port):$(port) -p=4040:4040 -v=$(current_dir):$(target_dir) -w=$(target_dir) jupyter/pyspark-notebook jupyter lab --no-browser --notebook-dir=$(target_dir) --allow-root --port=$(port) --ip=$(ip)

kaggle:
	docker run --rm -it --name=kaggle --hostname=localhost -p=$(port):$(port) -v=$(current_dir):$(target_dir) -w=$(target_dir) kaggle/python jupyter lab --no-browser --notebook-dir=$(target_dir) --allow-root --port=$(port) --ip=$(ip)

tensor:
	docker run --rm -it --name=tensorflow -p=$(port):$(port) -v=$(current_dir):$(target_dir) -w=$(target_dir) tensorflow/tensorflow jupyter lab --no-browser --notebook-dir=$(target_dir) --allow-root --port=$(port) --ip=$(ip)

native:
	jupyter lab --allow-root --port=$(port) --ip=$(ip)

up-docker:
	docker build -t coinbase-client -f Dockerfile .
	docker rm -f coinbase-client || true
	docker run -d --restart always --name coinbase-client --hostname coinbase-client --net=host coinbase-client
