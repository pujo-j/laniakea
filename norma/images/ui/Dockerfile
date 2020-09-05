FROM node:10.16.0 as ui

ENV PREFECT_UI_COMMIT_ID=85da2fac893d41e494dba07d0b5c4bfa1e21b820

RUN mkdir -p /app
WORKDIR /app
RUN git clone https://github.com/PrefectHQ/ui.git
WORKDIR /app/ui/
RUN git checkout $PREFECT_UI_COMMIT_ID

# Configure env and patch
ENV VUE_APP_ENVIRONMENT=production
ENV VUE_APP_BACKEND=SERVER
ENV VUE_APP_SERVER_URL="/gql/graphql"

RUN rm .env
COPY 0001-First-patchset.patch .
RUN git apply 0001-First-patchset.patch

# Install dependencies
RUN npm install

# Build static files
RUN npm run build

FROM golang:1.15.0 as go
RUN mkdir -p /app
WORKDIR /app
COPY go.mod /app/
COPY imports.go /app/
RUN go mod vendor
COPY *.go /app/
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=go /app/ui /app/

COPY --from=ui /app/ui/dist/ /app/dist/

COPY --from=go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


WORKDIR /app

EXPOSE 8080
CMD ["/app/ui"]

