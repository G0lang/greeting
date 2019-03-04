# Build stage
FROM golang
ENV GO111MODULE=on
WORKDIR /src
COPY go.mod /src
COPY go.sum /src
RUN go mod download
COPY . /src
RUN SHORTHASH=$(git log -n1 --pretty="format:%h") && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags  " -w -s \
    -X main.VERSION=$SHORTHASH " -a -o greeting .

# Final stage
FROM gcr.io/distroless/base
COPY --from=0 /src/greeting /
EXPOSE 8080
ENTRYPOINT ["/greeting"]