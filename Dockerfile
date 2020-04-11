FROM golang:alpine
WORKDIR /go/src/Hackweek2020
COPY main /
ENTRYPOINT ["/main"]
