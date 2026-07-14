package main

import "fmt"

// Person is a simple struct with one field.
type Person struct {
	Name string
}

// Greet is a value-receiver method — gets a copy, can't mutate the original.
func (p Person) Greet() string {
	return "Hello, " + p.Name
}

// Employee embeds Person, promoting its fields/methods.
type Employee struct {
	Person
	Salary float64
}

// UpdateEmail-style mutation example via pointer receiver.
type User struct {
	Name  string
	Email string
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	// --- Zero values ---
	var u1 User
	fmt.Printf("zero-value User: %+v\n", u1) // Name:"" Email:""

	// --- Creating instances ---
	u2 := User{Name: "Abhishek", Email: "abhi@old.com"}
	fmt.Printf("u2 before update: %+v\n", u2)

	u2.UpdateEmail("abhi@new.com") // Go auto-takes &u2 here
	fmt.Printf("u2 after update:  %+v\n", u2)

	fmt.Println()

	// --- Embedding / composition ---
	e := Employee{
		Person: Person{Name: "Riya"},
		Salary: 90000,
	}

	// e.Name is "promoted" from the embedded Person struct
	fmt.Println("Employee name (promoted field):", e.Name)
	fmt.Println("Employee greet (promoted method):", e.Greet())
	fmt.Println("Employee salary:", e.Salary)

	fmt.Println()

	// --- Value receiver copy behavior ---
	p1 := Person{Name: "Original"}
	greetingCopy := p1 // this is a full copy of the struct
	greetingCopy.Name = "Changed copy"

	fmt.Println("p1.Name (untouched):", p1.Name)
	fmt.Println("greetingCopy.Name:", greetingCopy.Name)
}
