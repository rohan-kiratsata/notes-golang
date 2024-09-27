build go programs

```
// GO envs
go env

// build for specific OS
// no need to specify the OS if you are on that OS
GOOS="linux" go build
GOOS="windows" go build
GOOS="darwin" go build

// build for specific architecture
GOARCH="amd64" go build
GOARCH="arm64" go build
```
