package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	firstTask(5)
	secondTask()
}

//первое задание
func firstTask(count int) {
	var (
		wg      = sync.WaitGroup{}
		counter int
	)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			counter += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

//второе задание
func secondTask() {
	var mutex sync.Mutex
	go func() {
		defer mutex.Unlock()
		mutex.Lock()
		fmt.Println("Text")
	}()
	time.Sleep(2 * time.Second)
}

type Set struct {
	sync.Mutex
	mm map[float64]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[float64]struct{}{},
	}
}

func (s *Set) Add(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type SetRw struct {
	sync.RWMutex
	mm map[float64]struct{}
}

func NewRwSet() *SetRw {
	return &SetRw{
		mm: map[float64]struct{}{},
	}
}

func (s *SetRw) AddRw(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *SetRw) HasRw(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}
