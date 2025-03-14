# powershell

init:
	go install github.com/swaggo/swag/cmd/swag@v1.8.7
	go install github.com/vektra/mockery/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/4meepo/tagalign/cmd/tagalign@latest

start:
	go run ./cmd/main.go serve

test:
	go test -v ./...

mock:
	mockery --all --dir=internal/core/port --output=generated/mocks --case=underscore

pb:
	protoc -I=proto --go_out=. --go-grpc_out=. proto/beef.proto

swag:
	swag fmt
	swag init -p snakecase -g ./cmd/protocol/rest.go -o ./docs

fmt:
	tagalign -fix -sort -order "gorm,json,url,query,body,form,validate,envconfig,default,example" -strict ./...
