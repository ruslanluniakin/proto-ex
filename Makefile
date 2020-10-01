
install-proto:
	go install github.com/golang/protobuf/protoc-gen-go

gen-proto:
	protoc \
		-I=. \
		-I=${GOPATH}/src \
		--go_out=plugins=grpc:. \
		api/pingpongapi/pingpongapi.proto