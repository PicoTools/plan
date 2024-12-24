BIN_DIR=$(PWD)/bin
PLAN_DIR=$(PWD)/cmd/plan
CC=gcc
CXX=g++
ANTLR_VERSION=4.13.2

.PHONY: plan
plan: gen-parser
	@mkdir -p ${BIN_DIR}
	@echo "Building PLAN..."
	@CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -o ${BIN_DIR}/plan ${PLAN_DIR}

.PHONY: test-plan
test-plan: plan
	@echo "Testing language..."
	@${BIN_DIR}/plan samples/assert.pico
	@echo "  assert.pico: OK"
	@${BIN_DIR}/plan samples/recursion_max_depth.pico 1>/dev/null 2>/dev/null || true
	@echo "  recursion_depth.pico: OK"
	@${BIN_DIR}/plan samples/cast.pico 1>/dev/null 2>/dev/null
	@echo "  cast.pico: OK"
	@${BIN_DIR}/plan samples/files.pico
	@echo "  files.pico: OK"

.PHONY: gen-parser
gen-parser: download-antlr
	@echo "Generating parser..."
	@java -jar files/antlr-${ANTLR_VERSION}.jar -o pkg/parser -visitor -package parser -Dlanguage=Go -Xexact-output-dir ./grammar/PLAN.g4

download-antlr:
	@if [ ! -f files/antlr-${ANTLR_VERSION}.jar ]; then echo "Downloading antlr ${ANTLR_VERSION}..." && curl -s -o files/antlr-${ANTLR_VERSION}.jar https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar; fi
