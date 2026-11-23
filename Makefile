.PHONY: check installdeps install test

check:
	go build -o /dev/null ./...

test:
	go test -v ./...
