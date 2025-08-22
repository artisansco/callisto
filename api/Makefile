# Include environment variables
include .env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## upgradeable: list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	@go run github.com/oligot/go-mod-upgrade@latest


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...


# Run the application
run: build
	@go run cmd/api/


# Build the application
all: build test


build:
	@echo "Building..."
	@go build -o ./bin/api cmd/api/main.go


## build/api: build the cmd/api application
.PHONY: build/all
build/all:
	@echo 'Building cmd/api...'
	go build -ldflags="-s" -o=./bin/api ./cmd/api
	@echo 'Building for linux...'
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/api_linux_amd64 ./cmd/api
	@echo 'Building for OSX/darwin...'
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s" -o=./bin/api_darwin_amd64 ./cmd/api
	@echo 'Building for windows...'
	GOOS=windows GOARCH=amd64 go build -ldflags="-s" -o=./bin/api_windows_amd64 ./cmd/api


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -rf ./bin

# Live Reload
watch:
	@if command -v air > /dev/null; then \
		air \
			--build.cmd "make build" --build.bin "./bin/api" --build.delay "100" \
			--build.exclude_dir "" \
			--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
			--misc.clean_on_exit "true"; \
		echo "Watching..."; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air \
				--build.cmd "make build" --build.bin "./bin/api" --build.delay "100" \
				--build.exclude_dir "" \
				--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
				--misc.clean_on_exit "true"; \
			echo "Watching..."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi


.PHONY: all build run test clean watch

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #

## db/sqlite: connect to database using sqlite
.PHONY: db/sqlite
db/sqlite:
	sqlite ${DATABASE_URL}

## db/migrations/new name=$1: creates a new database migration
.PHONY: db/migrations/new
migrations/new:
	@echo 'Creating migration files for ${name}'
	goose -dir=migrations create ${name} sql

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
migrations/up: confirm
	@echo 'Running up migrations...'
	goose -dir=migrations sqlite ${DATABASE_URL} up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down
migrations/down:
	@echo 'Running down migrations...'
	goose -dir=migrations sqlite ${DATABASE_URL} down
