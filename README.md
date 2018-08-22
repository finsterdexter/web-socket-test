# finsterdexter/web-socket-test
Docker-image for WebSocket connectivity test. This container exposes two
websockets, at `ws://localhost:8010/ws_echo` and
`ws://localhost:8010/ws_broadcast`.

`ws_echo` is a standard websocket test, i.e. you can send a message and it will
bounce it back to you.

`ws_broadcast` just sends a continuous stream of messages without any input from
the client.

## How to
### Start contaienr
```bash
$ docker run -p 8010:8010 --name web-socket-test finsterdexter/web-socket-test
```

### Connect to container with WebSocket
```bash
$ sudo npm install wscat
$ ./node_modules/ws/bin/wscat --connect ws://$(CONTAINER_HOST_IP_ADDRESS):8010/ws_echo

connected (press CTRL+C to quit)
> Hello world

< Server received from client: Hello world
>
```

## Original docker image
Created by ksdn117