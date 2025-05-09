#Builder
FROM golang:1.23-alpine3.20 AS build
WORKDIR /app
RUN apk update && apk add --no-cache gcc git musl-dev
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# RUN go get -u github.com/vektra/mockery/cmd/mockery 
RUN pwd
COPY . .
RUN go mod vendor
RUN go build -ldflags '-w -s' -a -o ./api ./cmd

# Distribution
FROM alpine:latest
RUN apk update && apk add --no-cache
COPY --from=build /app/api /app/api
COPY --from=build /app/config.toml /app/config.toml
RUN  chmod +x /app/*
EXPOSE 4000
CMD ["/app/api"]