GO_FILES ?= ./...
GO_UNIT ?= ./modules/...
GO_INTEGRATION ?= ./routes/...
GO = GO111MODULE=on go

install:
	@echo "[+] install"
	$(GO) get -v -t -d $(GO_FILES)
.PHONY: install

start: install
	@echo "[+] start"
	@go run main.go
.PHONY: start

dev: scripts/bin/watcher
	@echo "[+] start in development mode"
	@scripts/bin/watcher
.PHONY: watch

unit-tests:
	@echo "[+] run unit test"
	@$(GO) test -race $(GO_UNIT)
.PHONY: unit-tests

integration-tests:
	@echo "[+] run integration test"
	@$(GO) test -race $(GO_INTEGRATION)
.PHONY: integration-tests

test: unit-tests integration-tests
.PHONY: test

lint: vet golangci-lint revive sec
.PHONY: lint

revive: scripts/bin/revive
	@echo "lint via revive"
	@scripts/bin/revive \
		-formatter stylish \
		-exclude ./vendor/... \
		$(GO_FILES)
.PHONY: revive

golangci-lint: scripts/bin/golangci-lint
	@echo "lint via golangci-lint"
	@scripts/bin/golangci-lint run \
		--config ./scripts/configs/.golangci.yml \
		$(GO_FILES)
.PHONY: golangci-lint

sec: scripts/bin/gosec
	@echo "lint via gosec"
	@scripts/bin/gosec -quiet \
		-exclude=G104,G107,G108,G201,G202,G204,G301,G304,G401,G402,G501 \
		-conf=./scripts/configs/gosec.json \
		$(GO_FILES)
.PHONY: sec

vet:
	@echo "lint via go vet"
	@$(GO) vet $(GO_FILES)
.PHONY: vet

scripts/bin/watcher: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/watcher github.com/canthefason/go-watcher/cmd/watcher

scripts/bin/revive: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/revive github.com/mgechev/revive

scripts/bin/golangci-lint: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

scripts/bin/gosec: scripts/go.mod
	@cd scripts; \
	$(GO) build -o ./bin/gosec github.com/securego/gosec/cmd/gosec
