FROM golang:1.18.3-alpine3.15 as builder

LABEL author="enronchain"

RUN apk add --update-cache \
    git \
    gcc \
    musl-dev \
    linux-headers \
    make \
    wget

RUN git clone https://github.com/enronchain/enron.git /enron && \
    #chmod -R 755 /enron && \
    chmod -R 755 /enron
WORKDIR /enron
RUN make install

# final image
FROM golang:1.18.3-alpine3.15

RUN mkdir -p /data

VOLUME ["/data"]

COPY --from=builder /go/bin/enrond /usr/local/bin/enrond

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["enrond"]