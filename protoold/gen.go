package protoold

//go:generate protoc -I . --go_out=paths=source_relative:. --twirp_out=paths=source_relative:. api.proto
