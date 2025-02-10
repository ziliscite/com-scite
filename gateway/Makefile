.PHONY: gen-auth
gen-auth:
	@protoc \
		--proto_path=proto "proto/auth.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative
