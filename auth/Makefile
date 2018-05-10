all: auth.pb.go internal.pb.go

auth.pb.go: auth.proto
	protoc --go_out=plugins=grpc:. $?

internal.pb.go: internal.proto
	protoc --go_out=plugins=grpc:. $?
