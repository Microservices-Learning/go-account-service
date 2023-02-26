proto:
	rm -rf go
	mkdir go
#	protoc client/proto/**/*.proto --go_out=plugins=grpc:
	protoc --proto_path=grpc/proto grpc/proto/**/*.proto --go_out=go
	protoc --proto_path=grpc/proto grpc/proto/**/*.proto --go-grpc_out=./go
	ls go/*/*/*/*.pb.go | xargs -n1 -Ix bash -c 'sed s/,omitempty// x > x.tmp && mv x{.tmp,}'
	rm -rf grpc/gen-proto/
	cp -rf go/* ../
	rm -rf go
server:
	go run main.go
