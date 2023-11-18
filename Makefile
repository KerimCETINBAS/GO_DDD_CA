export GO_ENV=development


start\:dev:
	gow run ./...
test\:unit:
	bash -c "go test  ./internal/domain/... ./internal/application/...  -v"