all: vm.pb.go

vm.pb.go: vm.proto
	protoc --go_out=plugins=grpc:. $?
