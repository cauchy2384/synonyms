.PHONY: lint
lint:
	golangci-lint run -c .golangcilint.yml

.PHONY: test
test: 
	go test -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov	| grep total

build:
	CGO_ENABLED=0 go build -o ./app/synonyms ./cmd/synonyms/

build-image:
	docker build -t cauchy2384/synonyms:latest . -f ./deployment/Dockerfile