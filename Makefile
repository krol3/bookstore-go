default: run

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

secrets:
	gitleaks detect --source . -v

semgrep:
	semgrep --config=auto .

gosec:
	gosec  ./...

chain-bench:
	echo "Add chain bench"

semgrep:
	semgrep --config=auto .
