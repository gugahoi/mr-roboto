FROM golang:1.10 as builder

LABEL author="Gustavo Hoirisch"
LABEL email="docker@gustavo.com.au"

WORKDIR /go/src/gtihub.com/gugahoi/mr-roboto/

ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/mr-roboto -a -ldflags '-extldflags "-static"' .

FROM scratch
COPY --from=builder /build/mr-roboto /app
ENTRYPOINT [ "/app" ]
