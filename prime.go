package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuProfile = flag.String("cpu", "", "write cpu profile to `file`")
var memProfile = flag.String("memory", "", "write memory profile to `file`")

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
	flag.Parse()
	if *cpuProfile != "" {
		f, er := os.Create(*cpuProfile)
		if er != nil {
			fmt.Println("Error in creating file for writing cpu profile: ", er)
			return
		}
		defer f.Close()

		if e := pprof.StartCPUProfile(f); e != nil {
			fmt.Println("Error in starting CPU profile: ", e)
			return
		}
		defer pprof.StopCPUProfile()

	}

	if *memProfile != "" {
		f, er := os.Create(*memProfile)
		if er != nil {
			fmt.Println("Error in creating file for writing memory profile to: ", er)
			return
		}
		defer f.Close()
		runtime.GC()
		if e := pprof.WriteHeapProfile(f); e != nil {
			fmt.Println("Error in writing memory profile: ", e)
			return
		}
	}
	prime(100000)

}
