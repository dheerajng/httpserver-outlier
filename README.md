# HTTPServer-Outlier
This is a simple HTTP server running on port 4040. It sends its own IP address in response.

If a request is sent to _*/misbehave*_ path, then 503 Service Unavailable response is sent.


## Building the Server

### Build
Build container using below command

```
make build
```

### Clean
```
make clean
```

### Run HTTP server
```
make run
```

### Build and Run

```
make docker-build-run
```

## Testing

`httpserver-outlier` runs on port 4040 on localhost.

1. Send GET request on '/'. Response contains the external IP address of container.

```
$ curl http://localhost:4040/
Hi there, You asked for ! I am 172.17.0.2
$
```

2. Send GET request on `/misbehave` path. The response is 503 HTTP error with message "MISBEHAVING! HTTP status code returned from <IP-address>!"
```
$ curl http://localhost:4040/misbehave
MISBEHAVING! HTTP status code returned from 172.17.0.2!
```
