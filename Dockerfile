#
# Build image: docker build -t bianjieai/spartan .
#
FROM golang:1.18-alpine3.15 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

WORKDIR /spartan

# Add source files
COPY . .

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

# ----------------------------
FROM alpine:3.15

# Set up dependencies
ENV PACKAGES make gcc openssl

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

COPY --from=builder /spartan/build/ /usr/local/bin/