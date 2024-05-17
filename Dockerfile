FROM golang:1.22-alpine3.19 AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM deps AS build

RUN adduser -D -u 1001 ayumu

RUN apk update && apk add upx

COPY . .

# https://mt165.co.uk/blog/static-link-go/
RUN go build \
  -ldflags="-extldflags '-static' -s" \
  -o ayumu-bin ./cmd

RUN upx -5 ./ayumu-bin

FROM scratch

WORKDIR /

# Copy user
COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/ayumu-bin ayumu-bin

USER ayumu

CMD ["/ayumu-bin"]
