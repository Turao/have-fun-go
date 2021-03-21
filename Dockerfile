FROM golang:latest as base
WORKDIR /go-cards
# statically compile the binary
# (i.e. do not include links to C libraries in the resulting binary)
ENV CGO_ENABLED=0
# (leverage docker caching optimizations)
# copy package management related files first
# since these will change less frequently
COPY go.mod .
COPY go.sum .
RUN go mod download
# then we can copy the rest of the project...
COPY . .


FROM base as build
RUN go build -o ./go-cards


FROM base as unit-tests
RUN go test -v .

FROM golangci/golangci-lint:latest as lint-base
FROM base as lint
COPY --from=lint-base /usr/bin/golangci-lint /usr/bin/golangci-lint
RUN golangci-lint run --timeout 5m0s ./...

FROM scratch as deploy
WORKDIR /go-cards
COPY --from=build /go-cards/go-cards ./go-cards
CMD ["./go-cards"]