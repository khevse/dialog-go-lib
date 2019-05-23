.DEFAULT_GOAL=testall

PACKAGES_WITH_TESTS:=$(shell go list -f="{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}" ./... | grep -v '/vendor/' | grep -v '/kafka')
TEST_TARGETS:=$(foreach p,${PACKAGES_WITH_TESTS},test-$(p))
TEST_OUT_DIR:=testout

.PHONY: mod
mod:
	rm -rf vendor
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

.PHONY: static
static: mod
ifeq ($(shell command -v esc 2> /dev/null),)
	go get -u -v github.com/mjibson/esc
endif
	$(eval $@_target := github.com/dialogs/dialog-go-lib/db/migrations/test)
	rm -f $($@_target)/static.go
	go generate $($@_target)

.PHONY: mocks
mocks: static
ifeq ($(shell command -v mockery 2> /dev/null),)
	go get -u -v github.com/vektra/mockery/.../
endif
	$(eval $@_source := kafka)
	$(eval $@_target := ${$@_source}/mocks)
	rm -f $($@_target)/IReader.go
	rm -f $($@_target)/IWriter.go
	mockery -name=IReader -dir=${$@_source} -recursive=false -output=$($@_target)
	mockery -name=IWriter -dir=${$@_source} -recursive=false -output=$($@_target)

.PHONY: lint
lint: mocks
ifeq ($(shell command -v golangci-lint 2> /dev/null),)
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.14.1
endif
	golangci-lint run ./... --exclude "is deprecated"

.PHONY: testall
testall: lint
	rm -rf ${TEST_OUT_DIR}
	mkdir -p -m 755 $(TEST_OUT_DIR)
	$(MAKE) -j 5 $(TEST_TARGETS)
	@echo "=== tests: ok ==="

.PHONY: $(TEST_TARGETS)
$(TEST_TARGETS):
	$(eval $@_package := $(subst test-,,$@))
	$(eval $@_filename := $(subst /,_,$($@_package)))

	@echo "== test directory $($@_package) =="
	@GO111MODULE=on go test $($@_package) -v -race \
    -coverprofile $(TEST_OUT_DIR)/$($@_filename)_cover.out \
    >> $(TEST_OUT_DIR)/$($@_filename).out \
   || ( echo 'fail $($@_package)' && cat $(TEST_OUT_DIR)/$($@_filename).out; exit 1);
