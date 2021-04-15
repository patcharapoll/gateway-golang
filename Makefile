PROJECTNAME := $(shell basename "$(PWD)")
GOCMD = go
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64
GOBUILD = build

## bin: build go server to binary
.PHONY: build
build:
	env CGO_ENABLED=0 GOOS=${GOARCH} go build -a -installsuffix cgo -o bin/server cmd/server/main.go

.PHONY: dev
dev:
	go run cmd/server/main.go

.PHONY: watch
watch:
	CompileDaemon -include=Makefile --build="make build" --command=./bin/server --color=true --log-prefix=false

.PHONY: gqlgen
gqlgen:
	go run scripts/gqlgen.go --config internal/graph/gqlgen.yml