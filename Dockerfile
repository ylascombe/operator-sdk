FROM golang:1.9

RUN mkdir -p /go/src/github.com/operator-framework/operator-sdk
WORKDIR /go/src/github.com/operator-framework/operator-sdk
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN make dep
RUN make install
