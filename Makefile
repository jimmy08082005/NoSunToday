.PHONY: check installdeps install test

check:
	go build -o /dev/null ./...

installdeps:
	go mod download

test:
	go test -v ./...
