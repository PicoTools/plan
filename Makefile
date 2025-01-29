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

build: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor ${GOOS}/${GOARCH} ${VERSION}"
	@CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.${GOOS}.${GOARCH} ${PLAN_DIR}

darwin-arm64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor darwin/arm64 ${VERSION}"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.darwin.arm64 ${PLAN_DIR}

darwin-amd64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor darwin/amd64 ${VERSION}"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.darwin.amd64 ${PLAN_DIR}

linux-arm64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor linux/arm64 ${VERSION}"
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.linux.arm64 ${PLAN_DIR}

linux-amd64: gen-parser go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN runtime executor linux/amd64 ${VERSION}"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/plan.linux.amd64 ${PLAN_DIR}

go-lint:
	@echo "Linting Golang code..."
	@go fmt ${GOFILES}
	@go vet ${GOFILESNOTEST}

test-plan-base: build
	@echo "------------------------------"
	@echo "Testing language..."
	@${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/assert.pico
	@echo "  assert.pico: OK"
	@${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/recursion_max_depth.pico 1>/dev/null 2>/dev/null || true
	@echo "  recursion_depth.pico: OK"
	@${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/cast.pico 1>/dev/null 2>/dev/null
	@echo "  cast.pico: OK"
	@${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/files.pico
	@echo "  files.pico: OK"
	@echo "------------------------------"

test-plan-sorting: build
	@echo "---------------------"
	@echo "Testing gnome sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/gnome_sort.pico
	@echo "Testing bubble sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/bubble_sort.pico
	@echo "Testing shaker sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/shaker_sort.pico
	@echo "Testing selection sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/selection_sort.pico
	@echo "Testing insertion sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/insertion_sort.pico
	@echo "Testing heap sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/heap_sort.pico
	@echo "Testing merge sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/merge_sort.pico
	@echo "Testing comb sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/comb_sort.pico
	@echo "Testing shell sort"
	@time ${BIN_DIR}/plan.${GOOS}.${GOARCH} samples/sorting/shell_sort.pico
	@echo "---------------------"

gen-parser: download-antlr
	@echo "Generating parser..."
	@java -jar files/antlr-${ANTLR_VERSION}.jar -o pkg/parser -visitor -package parser -Dlanguage=Go -Xexact-output-dir ./grammar/PLAN.g4

download-antlr:
	@if [ ! -f files/antlr-${ANTLR_VERSION}.jar ]; then echo "Downloading antlr ${ANTLR_VERSION}..." && curl -s -o files/antlr-${ANTLR_VERSION}.jar https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar; fi
