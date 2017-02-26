.PHONY: all
all: build-deps build fmt vet lint test featuretest

GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell glide novendor)

APP_EXECUTABLE="out/battleship"

setup:
	curl https://glide.sh/get | sh
	go get -u github.com/golang/lint/golint
	go get github.com/DATA-DOG/godog/cmd/godog
	./scripts/pg_partman_install.sh

build-deps:
	glide install

update-deps:
	glide update

compile:
	mkdir -p out/
	go build -race -o $(APP_EXECUTABLE)

build: build-deps compile fmt vet lint

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	go test $(UNIT_TEST_PACKAGES) -p=1

test-cover-html: testdb.drop testdb.create testdb.migrate
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html
