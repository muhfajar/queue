test:
	go test ./... -v -cover

test-race:
	go test -race -coverprofile=coverage.txt -covermode=atomic