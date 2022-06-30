FROM golang:1.18.3 AS builder

COPY . /src
WORKDIR /src/app/center/admin

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/center/admin/bin /app

WORKDIR /app

EXPOSE 8001
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
