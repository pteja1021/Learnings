package main
import "fmt"

var m = make(map[int]int)

func main() {
	for _, n := range []int{5, 1, 9, 98, 6} {
		x := fib(n)
		fmt.Println(n, "fib", x)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	var f int
	if v, ok := m[n]; ok {
		f = v
	} else {
		f = fib(n-1) + fib(n-2)
		m[n] = f
	}

	return f
}