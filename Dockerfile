FROM golang:1.18 AS builder
RUN mkdir -p /app
ENV TZ=Asia/Jakarta
RUN GOCACHE=OFF
ARG GOPRIVATE_REPOS
RUN go env -w GOPRIVATE=$GOPRIVATE_REPOS
RUN go env -w GO111MODULE=on
ARG BITBUCKET_USER
ARG BITBUCKET_TOKEN
RUN git config --global url."https://$BITBUCKET_USER:$BITBUCKET_TOKEN@bitbucket.org/".insteadOf "https://bitbucket.org/"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY . /app
RUN cd /app && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/ven_baseapp_bo_backend/main ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/migrations/* migrations/
COPY --from=builder /go/bin/backend_test/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta
ENTRYPOINT ["/root/main"]
EXPOSE 8080
