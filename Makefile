# ==============================================================================
# Main
# ==============================================================================

run:
	go run ./main.go

build:
	go build ./main.go

test:
	go test -cover ./...

# ==============================================================================
# Protobuf commands
# ==============================================================================

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/*.proto

proto-clean:
	rm proto/*.pb.go

# ==============================================================================
# Docker compose commands
# ==============================================================================

develop:
	docker-compose -f docker-compose.yml up -d --build

develop-logs:
	docker-compose -f docker-compose.yml logs -f server

develop-down:
	docker-compose -f docker-compose.yml down

local:
	docker-compose -f docker-compose.local.yml up -d --build

local-logs:
	docker-compose -f docker-compose.local.yml logs -f server

local-down:
	docker-compose -f docker-compose.local.yml down
