package main

func main() {
	var sum int
	for i := 0; i < 20; i++ {
		go func(val int) {
			sum++
		}(i)
	}
}
