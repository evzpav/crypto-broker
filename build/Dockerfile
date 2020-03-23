# --- Base ----
FROM golang:1.14-stretch AS base
WORKDIR $GOPATH/src/github.com/evzpav/crypto-broker

# ---- Build ----
FROM base AS build
COPY . .
COPY config.yaml /go/bin/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/crypto-broker ./cmd/server/main.go

# --- Release ----
FROM alpine AS image
COPY --from=build /go/bin/crypto-broker /crypto-broker
COPY --from=build /go/bin/config.yaml .
ENTRYPOINT ["/crypto-broker"]
