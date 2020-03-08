FROM golang

RUN mkdir -p /go/src/PGraph

ADD . /go/src/PGraph

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run PGraph/client/cmd  -watch PGraph/client
