FROM golang:alpine as build

WORKDIR /go-cards

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./go-cards

FROM alpine:latest as deploy

WORKDIR /go-cards
COPY --from=build /go-cards/go-cards ./go-cards
CMD ["./go-cards"]