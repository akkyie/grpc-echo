.PHONY: all
all: build

.PHONY: build
build: build/grpc-echo

.PHONY: clean
clean:
	$(RM) build/grpc-echo

build/grpc-echo:
	go build -o $@ .

echo/%.pb.go: %.proto
	protoc --go_out=plugins=grpc:$(dir $@) $^

.PHONY: container
container: echo/echo.pb.go
	docker build -t grpc-echo .

.PHONY: deploy/gcr
deploy/gcr:
	docker tag grpc-echo gcr.io/aky-sh/grpc-echo
	docker push gcr.io/aky-sh/grpc-echo
