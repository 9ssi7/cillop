FROM golang:1.20-alpine AS builder

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

WORKDIR /

COPY go.* ./
RUN   --mount=type=cache,target=/go/pkg/mod \
  go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -o main ./cmd/main.go

FROM scratch

ENV PORT 8080

COPY --from=builder /main .
COPY --from=builder /locales ./locales
EXPOSE $PORT

CMD ["/main"]