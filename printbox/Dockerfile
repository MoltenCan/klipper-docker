FROM --platform=$BUILDPLATFORM golang:alpine as builder
COPY src/ /go/src
WORKDIR /go/src

# macos
RUN GOOS=linux GOARCH=arm64 go build -o /go/bin/linux/arm64/printbox .
# # rpi3+
RUN GOOS=linux GOARCH=arm GOARM=7 go build -o /go/bin/linux/arm/v7/printbox .
# # rpi
RUN GOOS=linux GOARCH=arm GOARM=6 go build -o /go/bin/linux/arm/v6/printbox .
# # linux
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/linux/amd64/printbox .

FROM --platform=$TARGETPLATFORM moltencan/docker-compose2
ARG TARGETPLATFORM
COPY --from=builder /go/bin/${TARGETPLATFORM}/printbox /usr/local/bin/printbox
COPY ep.sh /usr/local/ep.sh

ENTRYPOINT [ "/usr/local/ep.sh" ]
