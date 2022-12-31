.PHONY: default
default:
	@echo "..."

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o apoor-ssh-server ./cmd/server

.PHONY: vhs
vhs:
	docker-compose down && docker-compose up --build

