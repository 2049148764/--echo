# mutilstage 1 for compiling golang add zone datas
FROM golang:1.14-alpine as build-env
# All these steps will be cached
# Get dependancies - will also be cached if we won't change mod/sum
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN mkdir /app
WORKDIR /app
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .
RUN go mod download -x
# COPY the source code as the last step
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server

# App envs here
ENTRYPOINT ["/go/bin/server"]
