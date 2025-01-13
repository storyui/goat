.PHONY: build
build:
	@CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o goat
release:
	@goreleaser release --snapshot --clean