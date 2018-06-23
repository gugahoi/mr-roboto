FROM golang:1.10 as builder

LABEL author="Gustavo Hoirisch"
LABEL email="docker@gustavo.com.au"

RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/gtihub.com/gugahoi/mr-roboto/

ADD Gopkg.* ./
RUN dep ensure -vendor-only

ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/mr-roboto -a -ldflags '-extldflags "-static"' .

FROM scratch
COPY --from=builder /build/mr-roboto /app
COPY fixtures /
ENTRYPOINT [ "/app" ]
