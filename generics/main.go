package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cat := new(Cat)
	cat.Name = "meowius"

	catWriter := new(Writer[Cat])
	catWriter.Write(cat)

	car := new(Car)
	car.Model = "tesla"

	carWriter := new(Writer[Car])
	carWriter.Write(car)

	car2 := new(Car)
	car2.Model = "toyota"

	b, err := json.Marshal(car2)
	if err != nil {
		panic(err)
	}

	bw := new(BytesWriter)
	bw.Write(b)
}

type Car struct {
	Model string
}

type Cat struct {
	Name string
}

type Writer[I any] struct{}

func (w *Writer[I]) Write(s *I) (n int, err error) {
	b, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}

	return fmt.Println(string(b))
}

type BytesWriter struct{}

func (bw *BytesWriter) Write(b []byte) (n int, err error) {
	return fmt.Println(string(b))
}
