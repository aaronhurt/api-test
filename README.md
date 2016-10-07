## api-test

In one window ...

```
go run server/listen.go
```

In another window ...

```
CONSUL_HTTP_ADDR=localhost:9900 go run test.go
```
