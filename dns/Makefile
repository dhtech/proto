all: dns.pb.go enforcer.pb.go

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. $?
