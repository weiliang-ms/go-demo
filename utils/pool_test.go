package utils

import (
	"sync/atomic"
	"testing"
)

// WaitGroupPool测试用例

func TestWaitGroupPool(t *testing.T) {

	// 初始化10个协程
	wgp := NewWaitGroupPool(10)
	var total uint32

	// 模拟100个请求
	for i := 0; i < 100; i++ {
		// 每次加1
		wgp.Add()

		// 匿名函数
		go func(total *uint32) {
			defer wgp.Done()           // 执行加一后释放协程实例
			atomic.AddUint32(total, 1) // 每次加1
		}(&total)
	}
	// 等待处理
	wgp.Wait()

	if total != 100 {
		t.Fatalf("The size '%d' of the pool did not meet expectations.", total)
	}
}
