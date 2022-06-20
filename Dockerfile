FROM golang:1.17-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /have-we-met-before

## Run

FROM golang:1.17-bullseye AS run

WORKDIR /

COPY --from=build /have-we-met-before /have-we-met-before
COPY ./Template/index.html ./Template/index.html

ENTRYPOINT ["/have-we-met-before"]
