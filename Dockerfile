FROM golang:alpine as app-builder
WORKDIR /go/src/microserv-test
COPY . .
RUN apk add git
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=app-builder /go/bin/microserv /microserv
ENTRYPOINT ["/microserv"]
