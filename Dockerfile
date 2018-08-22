FROM golang:alpine as build
RUN apk add --no-cache \
	bash \
	git
COPY . /go/src/web-socket-test
WORKDIR /go/src/web-socket-test
RUN ["./build.sh"]

FROM scratch
COPY --from=build /go/src/web-socket-test/websocketd /
CMD ["/websocketd"]
EXPOSE 8010
