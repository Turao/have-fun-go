FROM golang:alpine as build

WORKDIR /go-cards
COPY . .

RUN go get -d -v .
RUN go install -v

FROM alpine:latest as deploy

WORKDIR /go-cards
COPY --from=build /go-cards/go-cards ./go-cards
CMD ["./go-cards"]