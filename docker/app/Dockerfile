# Build Environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /go-app

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment Environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /go-app/bin/app /go-app/

EXPOSE 6969
CMD ["/go-app/app"]

