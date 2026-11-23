.PHONY: check installdeps install

check:
	gofmt -e . > /dev/null
