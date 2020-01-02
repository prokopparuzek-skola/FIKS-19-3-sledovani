package main

import "fmt"
import "math"

const LIMIT = 50

func pairs(n uint) (p uint) {
	if n == 1 {
		return 0
	}
	fibs := make([]uint, n)
	fibs[0] = 1
	fibs[1] = 1
	for i := range fibs[2:] {
		fibs[i+2] = fibs[i+1] + fibs[i]
	}
	for i, fib := range fibs {
		for _, friend := range fibs[i+1:] {
			IsManFriend := friend%2 == 1
			IsManFib := fib%2 == 1
			if IsManFib == IsManFriend {
				continue
			}
			if isFib(fib + friend) {
				if gcd(fib, friend) != 1 {
					continue
				} else {
					p++
				}
			}
		}
	}
	return
}

func isFib(n uint) bool {
	var ques1, ques2 float64
	ques1 = math.Sqrt(5.0*float64(n)*float64(n) + 4.0)
	ques2 = math.Sqrt(5.0*float64(n)*float64(n) - 4.0)
	if ques1-float64(uint(ques1)) == 0 || ques2-float64(uint(ques2)) == 0 {
		return true
	}
	return false
}

func gcd(x uint, y uint) uint {
	if y > x {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		var M uint
		fmt.Scanf("%d", &M)
		if M > LIMIT {
			for j := i; j < N; j++ {
				fmt.Println(0)
			}
			return
		}
		fmt.Println(pairs(M))
	}
}
