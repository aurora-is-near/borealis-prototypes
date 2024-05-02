.PHONY: install-protoc-linux install-protoc-mac install-protoc-gen-go gen-proto

install-protoc-linux:
	apt install -y protobuf-compiler

install-protoc-mac:
	brew install protobuf

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

gen-proto-go:
	rm -f $(shell find go -name "*pb.go")
	protoc \
		--proto_path=types/proto/. \
		--go_out=paths=source_relative:go/. \
		--go-vtproto_out=paths=source_relative:go/. \
		$(shell find types/proto -name "*.proto")
