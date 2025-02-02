package main

import "fmt"

func maps() {
	ages := make(map[string]int)
	fmt.Println("1. Empty map:", ages)

	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println("2. After adding items:", ages)

	bobAge, exists := ages["Bob"]
	fmt.Printf("3. Bob's age: %d, exists: %v\n", bobAge, exists)

	ages["Bob"] = 26
	fmt.Println("4. After updating Bob's age:", ages)

	delete(ages, "Alice")
	fmt.Println("5. After deleting Alice:", ages)

	_, aliceExists := ages["Alice"]
	fmt.Printf("6. Does Alice still exist? %v\n", aliceExists)

	fmt.Printf("7. Length of map: %d\n", len(ages))

	fmt.Println("8. Iterating over map:")
	for name, age := range ages {
		fmt.Printf("   %s: %d\n", name, age)
	}

	fruits := map[string]int{
		"apple":  5,
		"banana": 2,
	}
	fmt.Println("9. Map literal:", fruits)

	combined := make(map[string]int)
	for k, v := range ages {
		combined[k] = v
	}
	for k, v := range fruits {
		combined[k] = v
	}
	fmt.Println("10. Merged map:", combined)

	for k := range combined {
		delete(combined, k)
	}
	fmt.Println("11. After clearing:", combined)
}
