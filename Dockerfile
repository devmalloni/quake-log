# build stage
FROM golang:alpine AS build-env
ADD . /go/src/quake-log
RUN apk add --no-cache gcc g++ make linux-headers git
RUN cd /go/src/quake-log && go build -o app

# final stage
FROM alpine
COPY --from=build-env /go/src/quake-log/app /usr/bin/quake-log
ENTRYPOINT ["quake-log"]
