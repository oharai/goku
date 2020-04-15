# goku
a full stack web server framework in Go

features

- [Routing](#Routing)

## Routing

### Basic Usage

```go
r := router.NewRouter()

r.GET("/", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
})

r.Run(":8080")
```

### Middleware

```go
r.GET("/", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}).Before(func (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("before"))
}).After(func (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("after"))
})
```