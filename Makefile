.PHONY: install-protoc-linux install-protoc-mac protoc-test

install-protoc-linux:
	apt install -y protobuf-compiler

install-protoc-mac:
	brew install protobuf

protoc-test:
	protoc \
		--proto_path=. \
		--descriptor_set_out=/dev/null \
		$(shell find . -name "*.proto")
