# Constructor Functions in Go

## Explanation

Go has **no constructors** as a language feature (no `__init__`, no `new ClassName()`). Instead, the community convention is to write a plain function, usually named `New` or `NewXxx`, that builds and returns a properly-initialized value.

```go
type User struct {
    Name  string
    Email string
    id    string // unexported — caller can't set this directly
}

func NewUser(name, email string) *User {
    return &User{
        Name:  name,
        Email: email,
        id:    generateID(),
    }
}
```

Callers do:

```go
u := NewUser("Abhishek", "abhi@example.com")
```

### Why bother, if you could just write `User{...}` directly?

1. **Enforce required setup** — e.g. generating an ID, setting defaults, opening a connection.
2. **Hide unexported fields** — external packages can't set `id` directly since it's lowercase; only `NewUser` (inside the same package) can.
3. **Validation** — a constructor can return an error if inputs are invalid:

```go
func NewUser(name, email string) (*User, error) {
    if email == "" {
        return nil, errors.New("email is required")
    }
    return &User{Name: name, Email: email}, nil
}

u, err := NewUser("Abhishek", "")
if err != nil {
    log.Fatal(err)
}
```

4. **Encapsulate complex initialization** — e.g. setting up a struct that wraps a mutex, a channel, or a map that must not be `nil`:

```go
type Cache struct {
    mu    sync.Mutex
    items map[string]string
}

func NewCache() *Cache {
    return &Cache{items: make(map[string]string)} // avoids nil map bugs
}
```

### The functional options pattern

For structs with many optional fields, Go favors **functional options** over long parameter lists or a Builder class:

```go
type Server struct {
    host string
    port int
    tls  bool
}

type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) { s.port = p }
}

func WithTLS() Option {
    return func(s *Server) { s.tls = true }
}

func NewServer(host string, opts ...Option) *Server {
    s := &Server{host: host, port: 8080} // defaults
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage:
srv := NewServer("localhost", WithPort(9090), WithTLS())
```

This scales cleanly as new optional settings are added, without breaking existing callers.

## Simplified

Go doesn't have a built-in "constructor" like other languages. Instead, people just write a normal function called `New...` that assembles the struct correctly — filling in defaults, hiding internal fields, checking for mistakes — and hands you back a ready-to-use value, the same way a factory hands you a fully assembled product instead of a box of loose parts.

## Diagram

```mermaid
flowchart LR
    A["Caller: NewUser(\"Abhishek\", \"abhi@example.com\")"] --> B["NewUser function"]
    B --> C{"email valid?"}
    C -- no --> D["return nil, error"]
    C -- yes --> E["generate id\nset defaults"]
    E --> F["return &User{...}, nil"]
    F --> A
