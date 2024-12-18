test-unit:
	go test -tags=unit -short -coverprofile=cp.out ./...
generate-mocks:
	mockgen -source beer/beer.go -destination beer/mocks/beer.go -package mocks
env:
	cp .env.example .env
run:
	go run cmd/api/main.go