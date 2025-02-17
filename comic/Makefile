.PHONY: gen-com
gen-com:
	@protoc \
		--proto_path=proto "proto/comic.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative
