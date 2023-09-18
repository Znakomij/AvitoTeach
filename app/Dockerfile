FROM debian:11-slim

RUN apt-get update && apt-get upgrade && \
    apt-get install wget -y && wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /usr/local/app
COPY source/go.mod source/go.sum ./
RUN go mod download
COPY /source/*.go ./
RUN go build -o app-go .

FROM debian:11-slim

COPY --from=builder /usr/local/app/app-go /

EXPOSE 8089:8089 6379:6379
CMD ["/app-go"]
