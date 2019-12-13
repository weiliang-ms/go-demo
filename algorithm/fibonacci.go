package main

import (
	"fmt"
	"time"
)

func main() {
	fibonacci(50)
}

func fibonacci(n int) []uint64 {
	start := time.Now()
	re := make([]uint64, n)
	re[0] = 1
	re[1] = 1
	for i := 2; i < n; i++ {
		re[i] = re[i-1] + re[i-2]

	}

	fmt.Printf("前%d个斐波那契数列为%v\n", n, re)
	fmt.Printf("花费时间为：%f秒\n", time.Since(start).Seconds())
	sum(re[n-3], re[n-2], re[n-1])
	return re
}

func sum(a, b, c uint64) {
	if a+b == c {
		fmt.Printf("%d + %d == %d", a, b, c)
	}
}
