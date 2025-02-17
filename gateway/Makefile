.PHONY: gen-auth
gen-auth:
	@protoc \
		--proto_path=proto "proto/auth.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative

.PHONY: gen-act
gen-act:
	@protoc \
		--proto_path=proto "proto/activation.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative

.PHONY: gen-com
gen-com:
	@protoc \
		--proto_path=proto "proto/comic.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative
