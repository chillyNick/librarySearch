run:
	go run cmd/grpc-server/main.go

generate-grpc:
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/library.proto

mock-repo:
	mockgen -destination=internal/mocks/repo.go -package=mocks github.com/chillyNick/librarySearch/internal/repo Repo

test:
	go test github.com/chillyNick/librarySearch/internal/...