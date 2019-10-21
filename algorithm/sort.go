package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num      = 100000
	rangeNum = 100000
)

/*
	[选择排序]
	从数组中选择最小元素，将它与数组的第一个元素交换位置。
	再从数组剩下的元素中选择出最小的元素，将它与数组的第二个元素交换位置
*/
func SelectSort(buf []int) {
	t := time.Now()
	times := 0 // 排序次数
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		// 调换值
		if min != i {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
	}
	fmt.Println("select sort times: ", times)
	fmt.Println("选择排序时间：", time.Since(t))
}

/*
	[冒泡排序]
	从左到右不断交换相邻逆序的元素，在一轮的循环之后，可以让未排序的最大元素上浮到右侧
*/
func BubbleSort(buf []int) {
	t := time.Now()
	times := 0 //排序次数
	for i := 0; i < len(buf)-1; i++ {
		flag := false // 假设未排序
		for j := 1; j < len(buf); j++ {
			if buf[j-1] > buf[j] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("BubbleSort times : ", times)
	fmt.Println("冒泡排序时间：", time.Since(t))
}

/*
	[快速排序]

*/
func FastSort(buf []int) {
	t := time.Now()
	fast(buf, 0, len(buf)-1)
	fmt.Println("快速排序时间：", time.Since(t))
}

func fast(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = key
	fast(a, l, i-1)
	fast(a, i+1, r)
}

func main() {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	var buf []int
	for i := 0; i < num; i++ {
		buf = append(buf, randSeed.Intn(rangeNum))
	}
	// 选择排序
	//SelectSort(buf)
	// 冒泡排序
	//BubbleSort(buf)
	// 快速排序
	FastSort(buf)
}
