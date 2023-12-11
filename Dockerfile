# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
RUN cd /go-ethereum && go mod download

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/bin/
RUN mkdir /data
VOLUME ["/data"]
EXPOSE 7545 7546 30303 30303/udp
ENTRYPOINT ["geth", "--ws", "--ws.addr", "0.0.0.0", "--http.addr", "0.0.0.0", "--http", "--defichain", "--http.corsdomain", "*", "--ws.origins", "*", "--http.api", "eth,web3,personal", "--ws.api", "eth,web3,personal", "--datadir", "/data"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
