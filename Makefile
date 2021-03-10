BUILD=build

.PHONY: release
release:
	@echo "Release v$(version)"
	@git pull
	@git checkout master
	@git pull
	@git checkout develop
	@git flow release start $(version)
	@git flow release finish $(version) -p -m "Release v$(version)"
	@git checkout develop
	@echo "Release v$(version) finished."

.PHONY: all
all: coverage.out

coverage.out: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -race -cover -covermode=atomic -coverprofile ./coverage.out.tmp ./...
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp

.PHONY: test
test: coverage.out

.PHONY: cover
cover: coverage.out
	@echo ""
	@go tool cover -func ./coverage.out

.PHONY: cover-html
cover-html: coverage.out
	@go tool cover -html=./coverage.out

.PHONY: benchmark
benchmark:
	@go test -bench=. ./...

.PHONY: clean
clean:
	@rm ./coverage.out
	@go clean -i ./...

.PHONY: generate
generate:
	@CGO_ENABLED=0 go generate ./...

.PHONY: lint
lint:
	@CGO_ENABLED=0 golangci-lint run ./...

${BUILD}/plugin/%.so: plugin/%/*.go
	@echo "Building $* plugin..."
	@go build -buildmode=plugin -o $@ ./plugin/$*

plugins: ${BUILD}/plugin/healthcheck.so

.PHONY: docs
docs: docs/units.md

docs/units.md: resources/unit-categories.json
	@echo "Generate $@..."
	@go run cmd/generate-docs/main.go -- $@

pkg/hyperdash/config/enum.go: resources/unit-categories.json
	@echo "Generate $@..."
	@go run cmd/generate-enum/main.go -- $@

www/scripts/dist/index.js: www/scripts/src/index.ts www/scripts/src/* www/scripts/package.json www/scripts/package-lock.json
	@echo "Generate js..."
	@cd ./www/scripts/; \
		./node_modules/.bin/webpack

pkg/hyperdash/ui/scripts/app.js: www/scripts/dist/index.js
	@cp $< $@

${BUILD}/hyperdash: $(shell find . -type f -print | grep -v vendor | grep "\.go") pkg/hyperdash/ui/scripts/app.js pkg/hyperdash/config/enum.go
	@echo "Building hyperdash..."
	@go generate ./cmd/hyperdash/
	@go build -o $@ ./cmd/hyperdash/

build: ${BUILD}/hyperdash plugins

run: ${BUILD}/hyperdash
	@$< run demo/angell.hcl


docker-build: $(shell find . -type f -print | grep -v vendor | grep "\.go") pkg/hyperdash/config/enum.go
	@echo "Building hyperdash..."
	@go generate ./cmd/hyperdash/
	@go build -o hyperdash ./cmd/hyperdash/
