package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func prime(lim int) []int {
	primes := make([]int, 0)
	for i := 1; i < lim; i++ {
		flag := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	defer profile.Start().Stop()
	x := prime(100000)
	fmt.Println(x)
}
