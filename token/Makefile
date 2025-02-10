.PHONY: gen-tok
gen-tok:
	@protoc \
		--proto_path=proto "proto/activation.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative

.PHONY: gen-auth
gen-auth:
	@protoc \
		--proto_path=proto "proto/auth.proto" \
		--go_out=pkg/protobuf --go_opt=paths=source_relative \
		--go-grpc_out=pkg/protobuf --go-grpc_opt=paths=source_relative

## migrate/new name=$1: create a new database migration
.PHONY: migrate/new
migrate/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext .sql -dir ./migrations ${name}
