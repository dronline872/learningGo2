/*
Программа вызова паники неявно и ее обработка
*/
package panic_program

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	time string
	text string
}

//Функция для передачи времени в ошибку
func NewErrorWithTime(text string) error {
	timeNow := time.Now().Format("15:04:05")
	return &ErrorWithTime{
		time: timeNow,
		text: text,
	}
}

//Форматирование ошибки
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime: %s", e.text, e.time)
}

//Функция для вызова паники
func PanicFunction() {
	defer func() {
		if v := recover(); v != nil {
			err := NewErrorWithTime(fmt.Sprint(v))
			fmt.Println(err)
		}
	}()

	numbers := []int{1, 2, 3}
	fmt.Println(numbers[6])
}
