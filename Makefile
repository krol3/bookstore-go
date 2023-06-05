test:
	go test -v ./...

test-coverage:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run: 
	go run main.go

format:
	gofmt -s -w . && git diff --exit-code

vet:
	go vet ./...

ci:
	act -P ubuntu-20.04=ghcr.io/catthehacker/ubuntu:act-20.04

go-mod:
	go mod download
	go mod verify

vulns:
	govulncheck ./...

ci-job:
	@echo $(arg)
	@act -P ubuntu-20.04=ghcr.io/catthehacker/ubuntu:act-20.04 -j $(arg)