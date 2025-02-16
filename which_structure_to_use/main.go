package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// filterUniquesUsingMap()
	// ipMapping()
	mappingPorts()
}

func filterUniquesUsingMap() {
	list := []string{"anime", "neovim", "youtube", "icecream", "anime", "anime", "anime", "anime"}

	uniqueFiler := make(map[string]struct{})

	for _, v := range list {
		uniqueFiler[v] = struct{}{}
	}

	fmt.Println(uniqueFiler)
}

func ipMapping() {
	incrementingIPCounts(10, 0, 200, 24)
	incrementingIPCounts(10, 0, 200, 25)
	incrementingIPCounts(10, 0, 200, 26)
	incrementingIPCounts(10, 0, 200, 28)
	for {
		fmt.Println(ipMap[10][0][200][24])
		time.Sleep(1 * time.Second)
		PrintMemUsage()
	}
}

type ipInfo struct {
	Request int
	Date    time.Time
}

var ipMap = make([][][][]*ipInfo, 255)

func incrementingIPCounts(oct1, oct2, oct3, oct4 int) {
	o1 := ipMap[oct1]
	if o1 == nil {
		ipMap[oct1] = make([][][]*ipInfo, 255)
	}
	o2 := ipMap[oct1][oct2]
	if o2 == nil {
		ipMap[oct1][oct2] = make([][]*ipInfo, 255)
	}
	o3 := ipMap[oct1][oct2][oct3]
	if o3 == nil {
		ipMap[oct1][oct2][oct3] = make([]*ipInfo, 255)
	}
	o4 := ipMap[oct1][oct2][oct3][oct4]
	if o4 == nil {
		ipMap[oct1][oct2][oct3][oct4] = &ipInfo{
			Request: 1,
			Date:    time.Now(),
		}
	} else {
		o4.Request++
	}
}

var portMap [math.MaxUint16]*ipInfo

func mappingPorts() {
	portMap[2000] = &ipInfo{
		Request: 1,
	}
	fmt.Println(portMap[2000])
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Current = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotal(over time) = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tGC count = %v\n", m.NumGC)
}
