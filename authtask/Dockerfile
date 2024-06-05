FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx
FROM --platform=$BUILDPLATFORM golang:alpine

RUN apk add clang lld

COPY --from=xx / /

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETPLATFORM

RUN xx-apk add musl-dev gcc

ENV CGO_ENABLED=1

RUN xx-go build -o motibase ./cmd/main.go && \
    xx-verify motibase

CMD ["./motibase"]    