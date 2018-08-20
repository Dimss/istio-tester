version=10
build:
	docker run -v $(PWD):/go/src/istio-tester -w /go/src/istio-tester -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 golang go build -o istio-tester
docker-build:
	docker build -t dimssss/istio-tester:$(version) .
docker-push:
	docker push dimssss/istio-tester:$(version)

