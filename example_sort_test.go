package gsort_test

import (
	"fmt"

	"github.com/dinalt/gsort"
)

// This example demonstrates usage of Sort function.
func ExampleSort() {
	type user struct {
		name string
		age  int
	}

	users := []user{
		{"John", 40},
		{"Bob", 23},
		{"Alice", 30},
	}

	gsort.Sort(users, func(a, b user) bool { return a.age < b.age })

  fmt.Printf("Sorted users: %v", users)
  // Output:
  // Sorted users: [{Bob 23} {Alice 30} {John 40}]
}
