# Structs in Go

## Explanation

A **struct** is Go's way of grouping related fields into one custom type. Go has no classes — structs plus methods and interfaces do the job classes do in other languages.

```go
type User struct {
    Name  string
    Email string
    Age   int
}
```

### Creating instances

```go
u1 := User{Name: "Abhishek", Email: "abhi@example.com", Age: 22}
u2 := User{}                 // zero value: "", "", 0
u3 := User{Name: "Riya"}     // Email and Age default to zero values
```

Every field has a **zero value** if not set: `""` for strings, `0` for numbers, `false` for bools, `nil` for pointers/slices/maps.

### Methods on structs

A method is a function with a **receiver** — it's how Go attaches behavior to a type:

```go
func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

`u` here is the receiver — calling `u1.Greet()` runs this method with `u1` bound to `u`.

There are two kinds of receivers:

- **Value receiver** `(u User)` — gets a *copy* of the struct. Changes inside the method don't affect the original.
- **Pointer receiver** `(u *User)` — gets a reference to the original. Changes persist.

```go
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail // modifies the real struct
}
```

Rule of thumb: if a method needs to mutate the struct, or the struct is large (avoid copying), use a pointer receiver.

### Struct embedding (composition, not inheritance)

Go achieves code reuse via **embedding** rather than class inheritance:

```go
type Person struct {
    Name string
}

type Employee struct {
    Person   // embedded — no field name, just the type
    Salary   float64
}

e := Employee{Person: Person{Name: "Abhishek"}, Salary: 90000}
fmt.Println(e.Name) // promoted field — accessible directly
```

`Employee` doesn't inherit from `Person` in the OOP sense — it simply contains a `Person` field, and Go "promotes" that field's members so you can access them as if they were flattened onto `Employee`.

### Struct tags

Structs can carry metadata used by libraries like `encoding/json`:

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}
```

## Simplified

A struct is a **custom labeled box** that bundles several related values together, like a form with fields: Name, Email, Age. Methods are actions you can perform on that box. If the method just *reads* the box, give it a copy (value receiver); if it needs to *write into* the real box, hand it the box's address (pointer receiver). Embedding is like putting one box inside another and being able to reach into the inner box's labels directly, without needing extra steps.

## Diagram

```mermaid
classDiagram
    class Person {
        +Name string
    }
    class Employee {
        +Person Person
        +Salary float64
    }
    Employee o-- Person : embeds

    class GreetMethod["u.Greet() -- value receiver"]
    class UpdateMethod["u.UpdateEmail() -- pointer receiver, mutates original"]
