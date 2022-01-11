FROM --platform=$BUILDPLATFORM golang:alpine as builder
COPY src/ /go/src
WORKDIR /go/src

# macos
RUN GOOS=linux GOARCH=arm64 go build -o /go/bin/linux/arm64/printbox .
# RUN go build -o /go/bin/linux/arm64/start .
# # rpi3+
# RUN GOOS=linux GOARCH=arm GOARM=7 go build -o /go/bin/linux/arm/v7/start .
# # rpi
# RUN GOOS=linux GOARCH=arm GOARM=6 go build -o /go/bin/linux/arm/v6/start .
# # linux
# RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/linux/amd64/start .

FROM --platform=$TARGETPLATFORM moltencan/docker-compose2
ARG TARGETPLATFORM
COPY --from=builder /go/bin/${TARGETPLATFORM}/printbox /usr/local/bin/printbox
COPY ep.sh /usr/local/ep.sh

ENTRYPOINT [ "/usr/local/ep.sh" ]
