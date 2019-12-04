FROM golang:alpine as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o powerstore_exporter .
FROM scratch
COPY --from=builder /build/powerstore_exporter /app/
WORKDIR /app
CMD ["./powerstore_exporter"]
