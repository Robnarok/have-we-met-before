FROM golang:1.17-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /have-we-met-before

## Run

FROM alpine:latest

WORKDIR /

COPY --from=build /have-we-met-before /have-we-met-before

ENTRYPOINT ["/have-we-met-before"]
