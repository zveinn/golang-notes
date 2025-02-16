//go:build linux

package main

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
)

func printOS() {
	fmt.Println("this is from linux")
	debug.PrintStack()
	stack := debug.Stack()
	fmt.Println("---------")
	stackLines := bytes.Split(stack, []byte{10})
	stackLines = stackLines[3:]
	for i, v := range stackLines {
		fmt.Println(i, string(v))
	}
	fmt.Println("---------")

	runtimeStack := make([]byte, 50000)
	n := runtime.Stack(runtimeStack, false)
	fmt.Println(string(runtimeStack[:n]))
}
