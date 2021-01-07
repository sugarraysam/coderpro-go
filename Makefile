TARGETS := test help
.PHONY: $(TARGETS)

test:
	@go test -covermode=count -coverprofile=.coverage.out ./challenges/...
	@go tool cover -func=.coverage.out
	@go test -race ./... >/dev/null 2>&1

help:
	@echo "make [ $(TARGETS) ]"
