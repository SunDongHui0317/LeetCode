package swordoffer

import (
	"sort"
	"sync"
)

func printNumbers(n int) []int {
	max := quickPow(10, n)
	list := make([]int, 0, max)
	var wg sync.WaitGroup
	wg.Add(max - 1)
	ch := make(chan int)

	for i := 1; i < max; i++ {
		go func(v int) {
			ch <- v
		}(i)
		go func() {
			defer wg.Done()
			list = append(list, <-ch)
		}()
	}

	wg.Wait()
	sort.Ints(list)
	return list
}

func quickPow(x int, n int) int {
	ans := 1
	for n > 0 {
		if d := n % 2; d == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}
