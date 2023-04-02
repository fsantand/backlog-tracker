build:
	@go build ./cmd/backlogtracker/main.go
run: build
	./main
