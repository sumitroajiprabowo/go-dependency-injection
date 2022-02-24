package example

type Person struct {
	Name string
}

type Customer struct {
	*Person
}

// Create provider functions for each of the dependencies customer with parameters person
func NewCustomer() *Customer {
	return &Customer{
		Person: &Person{
			Name: "John",
		},
	}
}
