package main

import (
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var mutex sync.Mutex
	var sum int
	for i := 0; i < 20; i++ {
		go func(val int) {
			defer mutex.Unlock()
			mutex.Lock()
			sum++
		}(i)
	}
}
