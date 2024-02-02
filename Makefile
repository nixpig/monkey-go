.PHONY: tidy
tidy: 
	go fmt ./...
	go mod tidy -v

.PHONY: test
test: 
	gotestsum --format testname ./...

.PHONY: test_coverage
test_coverage:
	gotestsum -- -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

