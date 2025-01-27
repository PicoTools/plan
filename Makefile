BIN_DIR=$(PWD)/bin
PLAN_DIR=$(PWD)/cmd/plan
CC=gcc
CXX=g++
ANTLR_VERSION=4.13.2

plan: gen-parser
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN..."
	@CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -o ${BIN_DIR}/plan ${PLAN_DIR}

test-plan-base: plan
	@echo "------------------------------"
	@echo "Testing language..."
	@${BIN_DIR}/plan samples/assert.pico
	@echo "  assert.pico: OK"
	@${BIN_DIR}/plan samples/recursion_max_depth.pico 1>/dev/null 2>/dev/null || true
	@echo "  recursion_depth.pico: OK"
	@${BIN_DIR}/plan samples/cast.pico 1>/dev/null 2>/dev/null
	@echo "  cast.pico: OK"
	@${BIN_DIR}/plan samples/files.pico
	@echo "  files.pico: OK"
	@echo "------------------------------"

test-plan-sorting: plan
	@echo "------------------------------"
	@echo "Testing bubble sort (3333 items)"
	@time ${BIN_DIR}/plan samples/sorting/bubble_sort.pico
	@echo "Testing shaker sort (3333 items)"
	@time ${BIN_DIR}/plan samples/sorting/shaker_sort.pico
	@echo "Testing comb sort (3333 items)"
	@time ${BIN_DIR}/plan samples/sorting/comb_sort.pico
	@echo "Testing insertion sort (3333 items)"
	@time ${BIN_DIR}/plan samples/sorting/insertion_sort.pico
	@echo "Testing shell sort (3333 items)"
	@time ${BIN_DIR}/plan samples/sorting/shell_sort.pico
	@echo "------------------------------"

gen-parser: download-antlr
	@echo "Generating parser..."
	@java -jar files/antlr-${ANTLR_VERSION}.jar -o pkg/parser -visitor -package parser -Dlanguage=Go -Xexact-output-dir ./grammar/PLAN.g4

download-antlr:
	@if [ ! -f files/antlr-${ANTLR_VERSION}.jar ]; then echo "Downloading antlr ${ANTLR_VERSION}..." && curl -s -o files/antlr-${ANTLR_VERSION}.jar https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar; fi
