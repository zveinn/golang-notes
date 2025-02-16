package main

import (
	"runtime"
	"runtime/debug"
)

// https://github.com/golang/go/blob/master/src/internal/syslist/syslist.go
// go build GOARCH=amd64 GOOS=windows .

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(4)
	// turn off garbage collection
	debug.SetGCPercent(-1)
	// 10 MB memory limit
	debug.SetMemoryLimit(10000000)
	// manually run garbage collection
	runtime.GC()

	printOS()
}
