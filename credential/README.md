go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --proto_path=. --go_out=. --go_opt=paths=source_relative credential.proto