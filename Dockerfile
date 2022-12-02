# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o /rand-server

## Deploy
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /rand-server /rand-server
EXPOSE 8080
USER nonroot:nonroot
CMD [ "/rand-server" ]
