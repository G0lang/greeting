include .env
SHORTHASH	:= $$(git log -n1 --pretty="format:%h")
CGO_ENABLED	:= 0
GOOS		:= linux
GOARCH		:= amd64
GO111MODULE	:= on
IMGNAME		:= greeting
IMGTAG		:= ${IMGNAME}:${SHORTHASH}
LATEST		:= ${IMGNAME}:latest

# Run commands with the debugger. (default: false)
DEBUG ?= false

# Show this help prompt.
help:
	@ echo
	@ echo '  Usage:'
	@ echo ''
	@ echo '    make <target> [flags...]'
	@ echo ''
	@ echo '  Targets:'
	@ echo ''
	@ awk '/^#/{ comment = substr($$0,3) } comment && /^[a-zA-Z][a-zA-Z0-9_-]+ ?:/{ print "   ", $$1, comment }' $(MAKEFILE_LIST) | column -t -s ':'
	@ echo ''
	@ echo '  Flags:'
	@ echo ''
	@ awk '/^#/{ comment = substr($$0,3) } comment && /^[a-zA-Z][a-zA-Z0-9_-]+ ?\?=/{ print "   ", $$1, $$2, comment }' $(MAKEFILE_LIST) | column -t -s '?=' | sort
	@ echo ''

# Show variable.
vars:
	@ echo '  Variable:'
	@ echo ''
	@ awk '/^[A-Z]+\t*[^\t]+?\:=/{ print "   ", $$1 }' $(MAKEFILE_LIST) | sort
	@ echo ''

# Build app 
build:
	@echo "Building Project Binary To ./bin"
	@GOARC=${GOARCH} GOOS=${GOOS} CGO_ENABLED=${CGO_ENABLED} go build -ldflags  " -w -s \
    -X main.VERSION=${SHORTHASH}" -a -o bin/greeting .

# Run app
run:
	@ go run -ldflags  " \
    -X main.VERSION=${SHORTHASH}" . 

# Run Test
coverage:
	@ go test ./... -cover

# Run Test
test:
	@ go test -v ./...

# Build docker image.
dbuild:
	@ docker build -t ${IMGTAG} .
	@ docker tag ${IMGTAG} ${LATEST}

# Build docker no cache
dbuild-nc:
	@ docker build --no-cache -t ${IMGTAG} .
	@ docker tag ${IMGTAG} ${LATEST}

# Run docker image.
drun:
	@ docker run --net=host --env-file=.env -p ${PORT}:${PORT} ${LATEST}

# Clean Docker.
dclean:
	@ docker container prune -f 
	@ docker image prune -f 
	@ docker image rm ${IMGTAG} ${LATEST}


# cockroach sql cli.
roachcli:
	@ docker exec -it roach1 ./cockroach sql --insecure

# cockroach run 
roachup:
	@ docker-compose up