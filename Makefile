.PHONY: all
all: build

.PHONY: build
build: build/grpc-echo

.PHONY: clean
clean:
	$(RM) build/grpc-echo

build/grpc-echo: echo/echo.pb.go
	go build -o $@ .

build/echo.pb: echo.proto
	protoc -o $@ $^

echo/%.pb.go: %.proto
	protoc --go_out=plugins=grpc:$(@D) $^

.PHONY: build/container
build/container: echo/echo.pb.go
	docker build -t grpc-echo .

.PHONY: deploy/gcr
deploy/gcr: build/container
	docker tag grpc-echo gcr.io/aky-sh/grpc-echo
	docker push gcr.io/aky-sh/grpc-echo

.PHONY: deploy/endpoints
deploy/endpoints: build/echo.pb api_config.yaml
	gcloud endpoints services deploy $^

.PHONY: deploy/gke
deploy/gke:
	kubectl apply -f gke/namespace.yaml -f gke

.PHONY: clean/gke
clean/gke:
	kubectl delete -f gke
