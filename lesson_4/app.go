package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	workers()
	TaskWithTimeout()
}

func workers() {
	var workers = make(chan struct{}, 1)
	i := 0
	for i != 1000 {
		fmt.Println(i)
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
			}()
			i++
		}()
	}
	fmt.Println(i)
}

func TaskWithTimeout() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM)
	for {
		go closeProgram(signalCh)
	}
}

func closeProgram(signalCh <-chan os.Signal) {
	_, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	if syscall.SIGTERM == <-signalCh {
		os.Exit(1)
	}
}
