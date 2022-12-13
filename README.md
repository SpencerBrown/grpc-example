This is a test/example for a Go client and server using grpc on a Mac.

# Setup

`brew install protobuf`
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
`go mod init github.com/SpencerBrown/grpc-example`
`go mod tidy`

# Generate gRPC Go sources

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpctest/api.proto`

