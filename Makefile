BIN_DIR=$(PWD)/bin
PLAN_DIR=$(PWD)/cmd/plan
CC=gcc
CXX=g++
GOFILES=`go list ./...`
GOFILESNOTEST=`go list ./... | grep -ve "test\|parser"` # ignore tests and parser (codegen)
ANTLR_VERSION=4.13.2
LDFLAGS=-ldflags="-s -w"
GOARCH=`go env -json | jq -r .GOARCH`
GOOS=`go env -json | jq -r .GOOS`
# colors
RED=\033[0;31m
GREEN=\033[0;32m
NC=\033[0m

build: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor for ${GOOS}/${GOARCH}"
	@CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.${GOOS}.${GOARCH} ${PLAN_DIR}

darwin-arm64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor for darwin/arm64"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.darwin.arm64 ${PLAN_DIR}

darwin-amd64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor for darwin/amd64"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.darwin.amd64 ${PLAN_DIR}

linux-arm64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor for linux/arm64"
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.linux.arm64 ${PLAN_DIR}

linux-amd64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor for linux/amd64"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.linux.amd64 ${PLAN_DIR}

go-lint:
	@go fmt ${GOFILES}
	@go vet ${GOFILESNOTEST}

test: test-pass test-fail test-sort

test-pass: build
	@echo "\nPASS tests\n"
	@for i in `find samples -type f -name "pass-*"`; do \
		printf "$$i: " && \
		${BIN_DIR}/plan.${GOOS}.${GOARCH} $$i 1>/dev/null 2>/dev/null ; \
		[ $$? -eq 0 ] && echo "${GREEN}OK${NC}" || echo "${RED}FAILED${NC}" \
		; \
	done

test-fail: build
	@echo "\nFAIL tests\n"
	@for i in `find samples -type f -name "fail-*"`; do \
		printf "$$i: " && \
		${BIN_DIR}/plan.${GOOS}.${GOARCH} $$i 1>/dev/null 2>/dev/null ; \
		[ $$? -eq 0 ] && printf "${RED}FAILED${NC}" || echo "${GREEN}OK${NC}" \
		; \
	done

test-sort: build
	@echo "\nSORT tests\n"
	@for i in `find samples/sorting -type f -name "*.pico"`; do \
		printf "$$i: " && \
		${BIN_DIR}/plan.${GOOS}.${GOARCH} $$i 1>/dev/null 2>/dev/null ; \
		[ $$? -eq 0 ] && echo "${GREEN}OK${NC}" || echo "${RED}FAILED${NC}" \
		; \
	done

gen-parser: download-antlr
	@java -jar files/antlr-${ANTLR_VERSION}.jar -o pkg/parser -visitor -package parser -Dlanguage=Go -Xexact-output-dir ./grammar/PLAN.g4

download-antlr:
	@if [ ! -f files/antlr-${ANTLR_VERSION}.jar ]; then echo "Downloading antlr ${ANTLR_VERSION}..." && curl -s -o files/antlr-${ANTLR_VERSION}.jar https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar; fi
