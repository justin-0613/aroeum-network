# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Garo in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-aroeum/
COPY go.sum /go-aroeum/
RUN cd /go-aroeum && go mod download

ADD . /go-aroeum
RUN cd /go-aroeum && go run build/ci.go install -static ./cmd/garo

# Pull Garo into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-aroeum/build/bin/garo /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["garo"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
