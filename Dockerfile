FROM golang:1.14 as builder

COPY ./ /opt/app
WORKDIR /opt/app 

RUN CGO_ENABLED=1 go build -race utils/blockchain_loader_cyberway/blockchain_loader_cyberway.go \
    && CGO_ENABLED=1 go build -race utils/scan_blockchain_cyberway/scan_blockchain_cyberway.go
RUN go build

FROM debian:10

COPY --from=builder /opt/app/stihi-backend /usr/local/bin/
COPY --from=builder /opt/app/blockchain_loader_cyberway /usr/local/bin/
COPY --from=builder /opt/app/scan_blockchain_cyberway /usr/local/bin/

RUN chmod +x /usr/local/bin/blockchain_loader_cyberway
RUN chmod +x /usr/local/bin/scan_blockchain_cyberway
RUN chmod +x /usr/local/bin/stihi-backend
