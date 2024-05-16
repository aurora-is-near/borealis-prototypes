.PHONY: install-protoc-linux install-protoc-mac protoc-test

install-protoc-linux:
	apt install -y protobuf-compiler

install-protoc-mac:
	brew install protobuf

protoc-test:
	protoc \
		--proto_path=types/proto/. \
		--descriptor_set_out=/dev/null \
		$(shell find types/proto -name "*.proto")
