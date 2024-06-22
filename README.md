# goerror

Core error interface in Go

### Install

```shell
go get github.com/prongbang/goerror
```

### How to use

- Unauthorized error

```go
err := goerror.NewUnauthorized()
```

- Other status in [status.go](status.go)
