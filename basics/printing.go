package main

import "fmt"

func printing() {
	fmt.Printf("Basic formatting: %d %.2f %s\n", 1, 3.14159, "pi")

	fmt.Printf("Field widths and alignment: |%5d| |%5.2f| |%-8s|\n", 123, 3.14159, "left")

	fmt.Printf("Flags: %+d %04d %#x\n", 1, 2, 17)

	var p *int
	fmt.Printf("Pointer, Boolean, String: %p %t %q\n", p, true, "Quote")

	fmt.Printf("Unicode and Escapes: %c \\n %U\n", 'α', 0x03B1)

	x := struct {
		Meow string
	}{
		Meow: "fdssf",
	}
	fmt.Printf("Print struct: %+v\n", x)

	fmt.Printf("Padding and truncating: %.5s %.5f\n", "Hello, World!", 3.14159265359)
}
