# http-server

Simple http-server to test routing + templating capabilities

### compiling templates

- ##### install qtc (quicktemplate compiler)
```
go get -u github.com/valyala/quicktemplate
go get -u github.com/valyala/quicktemplate/qtc
```
- ##### run qtc
```
~/go/bin/qtc
```

### building
```
go build [-o server]
```

### running
```
go build -o server
./server
```
*OR*
```
go run server.go
```
