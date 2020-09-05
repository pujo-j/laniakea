FROM golang:1.15.0 as go
RUN mkdir -p /app
WORKDIR /app
COPY * /app/

RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=go /app/wait /app/

WORKDIR /app

CMD ["/app/wait"]