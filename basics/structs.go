package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Define a struct
type Person struct {
	Name     string
	Age      int
	Email    string
	Birthday time.Time
}

// Method on struct (receiver)
func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Method to update age
func (p *Person) GrowOlder() {
	p.Age++
}

func structOperations() {
	person1 := Person{
		Name:     "Alice",
		Age:      30,
		Email:    "alice@example.com",
		Birthday: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println("1. Creating a struct:", person1)

	fmt.Printf("2. Accessing fields: Name: %s, Age: %d\n", person1.Name, person1.Age)

	person1.Age = 31
	fmt.Println("3. After update:", person1)

	fmt.Println("4. Using method:", person1.String())
	person1.GrowOlder()
	fmt.Println("4. After growing older:", person1)

	person2 := &Person{
		Name: "Bob",
		Age:  25,
	}
	fmt.Println("5. Using struct pointer:", person2)

	anon := struct {
		Field1 string
		Field2 int
	}{
		Field1: "Anonymous",
		Field2: 1,
	}
	fmt.Println("6. Anonymous struct:", anon)

	type Employee struct {
		Person
		Position string
	}

	employee := Employee{
		Person:   Person{Name: "Charlie", Age: 40},
		Position: "Developer",
	}
	fmt.Println("7. Struct composition:", employee)

	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	} else {
		fmt.Println("8. JSON marshal:", string(jsonData))
	}

	var person3 Person
	err = json.Unmarshal(jsonData, &person3)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	} else {
		fmt.Println("8. JSON unmarshal:", person3)
	}

	type User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	user := User{
		FirstName: "David",
		LastName:  "Smith",
	}
	userData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling user JSON:", err)
	} else {
		fmt.Println("9. Struct tags in JSON:", string(userData))
	}

	person4 := Person{Name: "Alice", Age: 31}
	fmt.Println("10. Struct equality:", person1 == person4)
}
