# 版本信息
GO_VERSION:=$(shell go env GOPATH)
GOPATH:=$(shell go env GOPATH)
GIT_BRANCH=$(shell git branch | grep '*' | awk '{print $$2}')
GIT_VERSION=$(shell git describe --tags --always)

.PHONY: info
info:
	@echo "==> GO_VERSION: $(GO_VERSION)"
	@echo "==> GOPATH: $(GOPATH)"
	@echo "==> GIT_BRANCH: $(GIT_BRANCH)"
	@echo "==> GIT_VERSION: $(GIT_VERSION)"

# make run service=xxx
.PHONY: run
run:
	go run ./cmd/main/... -conf=./configs

# make api service=xxx
.PHONY: api
api:
	go run ./cmd/proto/... -path=./api/${service}

# make config
.PHONY: config
config:
	go run ./cmd/proto/... -path=./internal/conf

# ping 请注意端口号
# 注意，请先运行服务
# make run service=admin-service
.PHONY: ping
ping:
	curl http://127.0.0.1:8081/api/v1/ping/hello && \
    echo "\n" && \
    curl http://127.0.0.1:8081/api/v1/ping/error && \
    echo "\n"
