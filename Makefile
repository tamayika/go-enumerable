.PHONY: test
test:
	go tool go2go test
	rm -rf *.go