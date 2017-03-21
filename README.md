commenTed

removes comments from text
useful to make templates in .go instead of .txt ...

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

------
Replacer

```go
func (u *User) FindByNAME(data string) error
// c:replace:up [NAME|ID] - [data string|id hide.Int64]
```

becomes

```go
func (u *User) FindByID(id hide.Int64) error
```
