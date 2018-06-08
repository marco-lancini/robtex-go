FROM golang:1.10.2-alpine3.7 AS builder

#WORKDIR /go/src/github.com/marco-lancini/robtex-go/
#COPY main.go /go/src/github.com/marco-lancini/robtex-go/


WORKDIR /go/src/robtex-go/
COPY . /go/src/robtex-go/


RUN CGO_ENABLED=0 go build -o /bin/robtex-go


FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /bin/robtex-go /bin/robtex-go
ENTRYPOINT ["/bin/robtex-go"]
