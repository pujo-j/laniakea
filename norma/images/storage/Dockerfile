FROM golang:1.15.0 as go
RUN mkdir -p /app
WORKDIR /app
COPY go.mod /app/
COPY imports.go /app/
RUN go mod vendor
COPY *.go /app/
RUN CGO_ENABLED=0 go build

FROM ubuntu as ubuntu

FROM scratch

COPY --from=go /app/storage /app/
COPY --from=ubuntu /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app

EXPOSE 8080

CMD ["/app/storage"]