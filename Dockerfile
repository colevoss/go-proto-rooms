FROM golang:latest AS build

# WORKDIR /build
WORKDIR $GOPATH/src/proto-test

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /build/main ./cmd/server/main.go
# RUN make build


FROM scratch

COPY --from=build /build/main /

EXPOSE 9000
ENTRYPOINT ["/main"]
