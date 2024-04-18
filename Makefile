build:
	go build -o bin/learn-golang-restful cmd/main.go

run:
	go build -o bin/learn-golang-restful cmd/main.go && ./bin/learn-golang-restful

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

mock:
	go generate ./...

test:
	go test -race -cover ./...