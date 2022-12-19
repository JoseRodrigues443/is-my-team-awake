.DEFAULT_GOAL := dev

.PHONY: dev
dev: ## dev build
dev: clean install generate vet fmt lint test mod-tidy build

.PHONY: ci
ci: ## CI build
ci: dev diff

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -f ./config/default.yaml
	rm -rf dist
	rm -f coverage.*
	rm -f ./logs/*.log
	rm -f ./db/*.db

.PHONY: install
install: env ## go install tools
	$(call print-target)
	cd tools && go install $(shell cd tools && go list -f '{{ join .Imports " " }}' -tags=tools)

.PHONY: generate
generate: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: vet
vet: ## go vet
	$(call print-target)
	go vet ./...


.PHONY: fmt
fmt: ## go fmt
	$(call print-target)
	go fmt ./...


.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	golangci-lint run

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out ./tests/
	go tool cover -html=coverage.out -o coverage.html

.PHONY: mod-tidy
mod-tidy: ## go mod tidy
	$(call print-target)
	go mod tidy
	cd tools && go mod tidy

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

.PHONY: build
build: ## goreleaser --snapshot --skip-publish --rm-dist
build: install
	$(call print-target)
	goreleaser --snapshot --skip-publish --rm-dist


.PHONY: release
release: ## goreleaser --rm-dist
release: install
	$(call print-target)
	goreleaser --rm-dist

.PHONY: run
run: ## go run
	@go run -race .

.PHONY: go-clean
go-clean: ## go clean build, test and modules caches
	$(call print-target)
	go clean -r -i -cache -testcache -modcache

.PHONY: env
env:
	npm i
	[ -e "./config/default.yaml" ] && echo "Env Exists" || cp ./config/example.default.yaml ./config/default.yaml


define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef


.PHONY: commit
commit:
	npm run commit

.PHONY: force-commit
force-commit:
	npm run force-commit

.PHONY: commit-all
commit-all:
	git add -A
	make commit

.PHONY: help
help:
	echo "Makefile command lists:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	echo "is-my-team-awake Help list:"
	go run . --help

.PHONY: alias
alias: ## Loads all the alias
	print "alias is-my-team-awake='go run .'"

.PHONY: demo
demo:
	sh docs/demo.sh
