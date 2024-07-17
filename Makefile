CURRENT_DIR = $(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/routes.go --output api/handler/docs	