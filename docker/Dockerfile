FROM golang:1.19
WORKDIR /usr/src/app
RUN git clone https://github.com/wintermute767/gRPCServerCalculator.git .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./cmd/main.go
CMD ["app"]