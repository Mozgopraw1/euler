package main

import "fmt"

func multiples(a, multi int) int {
	if a%multi == 0 {
		return 1
	}
	return 0
}

func main() {
	var sum, a, multi int
	multi = 3
	for i := 1; i <= 1000; i++ {
		a = i
		sum += multiples(a, multi)
	}
	fmt.Println(sum)
}
