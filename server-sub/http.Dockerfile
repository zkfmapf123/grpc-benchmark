ARG GO_VERSION=1.21.7
FROM golang:${GO_VERSION} AS builder 
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY http.go  .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap .

FROM arm64v8/alpine:latest AS Runner

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/bootstrap .

ENV DB_HOST=localhost
ENV DB_USER=root
ENV DB_PASSWORD=12341234
ENV DB_PORT=3306
ENV DB_DATABASE=Users
ENV PORT=3001

CMD ["./bootstrap"]




