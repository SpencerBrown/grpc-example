This is a test/example for a Go client and server using grpc on a Mac.

# Setup

`brew install protobuf`
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
`git clone https://github.com/SpencerBrown/grpc-example`

# initial setup if you are making changes or initializing the repo dependencies

most times you won't need to mess with the go module stuff

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpctest/api.proto`
`go mod init github.com/SpencerBrown/grpc-example`
`go mod tidy`

