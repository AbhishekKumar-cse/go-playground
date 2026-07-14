package main

import "fmt"

// Counter demonstrates why pointer receivers are needed for mutation.
type Counter struct {
	Count int
}

// Increment uses a pointer receiver so it mutates the real struct.
func (c *Counter) Increment() {
	c.Count++
}

// increment (plain function) shows the difference between value and pointer args.
func increment(n int) {
	n++ // only changes the local copy
}

func incrementPtr(n *int) {
	*n++ // changes the original, via its address
}

func main() {
	// --- Basic pointer mechanics ---
	x := 10
	p := &x // p holds the address of x

	fmt.Println("x:", x)
	fmt.Println("p (address):", p)
	fmt.Println("*p (value at address):", *p)

	*p = 20 // write through the pointer
	fmt.Println("x after *p = 20:", x)

	fmt.Println()

	// --- Pass by value vs pass by pointer ---
	a := 5
	increment(a)
	fmt.Println("after increment(a):", a) // still 5

	incrementPtr(&a)
	fmt.Println("after incrementPtr(&a):", a) // 6

	fmt.Println()

	// --- Pointer receivers on structs ---
	c := Counter{}
	c.Increment() // Go automatically takes &c here
	c.Increment()
	c.Increment()
	fmt.Println("Counter.Count after 3 increments:", c.Count) // 3

	// --- nil pointer safety ---
	var np *int
	if np == nil {
		fmt.Println("np is nil — dereferencing it would panic")
	}
}
