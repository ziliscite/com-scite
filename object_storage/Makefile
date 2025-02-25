.PHONY: gen-upl
gen-upl:
	@protoc \
		--proto_path=proto "proto/uploads.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative