FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
WORKDIR /build/zero
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./static /app/static
COPY ./config /app/config
RUN go build -ldflags="-s -w" -o /app/micro ./micro.go


FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/micro /app/micro
COPY --from=builder /app/static /app/static
COPY --from=builder /app/config  /app/config
CMD ["./micro", "-fig/config.yaml"]
