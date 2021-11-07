package main

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	time string
	text string
}

func NewError(text string) error {
	timeNow := time.Now().Format("15:04:05")
	return &ErrorWithTime{
		time: timeNow,
		text: text,
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime: %s", e.text, e.time)
}

func main() {
	panicFunction()
	fmt.Println("Программа работает")
}

func panicFunction() {
	defer func() {
		if v := recover(); v != nil {
			err := NewError(fmt.Sprint(v))
			fmt.Println(err)
		}
	}()

	numbers := []int{1, 2, 3}
	fmt.Println(numbers[6])
}
