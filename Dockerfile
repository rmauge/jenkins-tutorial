FROM golang:rc-alpine
EXPOSE 8080
ENV GOBIN=/app/
#ADD go.mod /go/src/webby/
ADD webby.go /go/src/webby/
WORKDIR /go/src/webby/
RUN go install ./webby.go
ENTRYPOINT ["/app/webby"]

