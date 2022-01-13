package panic_program

import "fmt"

func ExampleNewErrorWithTime() {
	result := NewErrorWithTime("test text")
	fmt.Println(result)
}
