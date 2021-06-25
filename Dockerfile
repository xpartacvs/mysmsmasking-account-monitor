FROM golang:1.16-alpine AS builder
WORKDIR /builder/src
COPY . .
RUN mkdir -p /builder/bin
RUN go build -ldflags="-s -w" -o /builder/bin/mysmsmasking-monitor main.go

FROM alpine:latest
LABEL maintainer="xpartacvs@gmail.com"
ENV TZ=Asia/Jakarta
ENV DISCORD_BOT_MESSAGE=Reminder\ akun\ MySMSMasking
ENV LOGMODE=disabled
ENV BALANCE_LIMIT=100000
ENV GRACE_PERIOD=7
ENV SCHEDULE=0\ 0\ *\ *\ *
WORKDIR /usr/local/bin
RUN apk update
RUN apk add --no-cache tzdata
COPY --from=builder /builder/bin/mysmsmasking-monitor .
CMD ["mysmsmasking-monitor"]
