commenTed

removes comments from text

```go
// c:remove
hello
// world
// c:end
```

becomes

```go
hello
world
```

------

```go
// c:remove
log.Println("hello") // c:too
// log.Println("world")
// c:end
```

becomes

```go
log.Println("world")
```
