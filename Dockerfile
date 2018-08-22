FROM golang:alpine
RUN apk add --no-cache \
	bash \
	git
COPY . /go/src/web-socket-test
WORKDIR /go/src/web-socket-test
RUN ["./build.sh"]
CMD ["./websocketd"]
EXPOSE 8010
