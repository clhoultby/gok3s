FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

COPY  . $GOPATH/src/gok3s

WORKDIR $GOPATH/src/gok3s

RUN go get
RUN go build -o /go/bin/gow3s


FROM scratch

COPY --from=builder /go/bin/gow3s /go/bin/gow3s

CMD [ "/go/bin/gow3s/main" ]
