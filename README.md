# Go Start

## Makefile

make {command}

### docker

- docker-build
- docker-up
- docker-restart
- docker-exec

### swagger

- swagger (update docs)

if you get an error
`export PATH=$(go env GOPATH)/bin:$PATH`

### unit test

- test
  
## Commands

### Test

```
go run cmd/cli/main.go test
```

## Swagger

```
http://{{domain}}/api/swagger
```

## Metrics

Basic Auth : system username and password from .env  

```
http://{{domain}}/api/system/metrics
```