package concurrent

import "sync"

// 需要先初始化一个缓冲通道，通过入队的阻塞来控制并发量
func GoWithLimitConCurrent(ch chan byte, f func()) {
	ch <- byte('t')
	go func() {
		f()
		<-ch
	}()
}

// GoN 同时启动多个协程，返回等待函数，适合多消费者场景
func GoN(n int, fn func(int)) func() {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			fn(i)
			wg.Done()
		}(i)
	}
	return wg.Wait
}

// 相当于java的栅栏，全部协程执行完成，主任务结束
// 需要已知任务数量
func ConcurrentWithRegularRoutings() {
	var wg sync.WaitGroup
	wg.Add(5) // 已知需要初始化几个协程的情况

	for i := 5; i > 0; i++ {
		go func() {
			wg.Done()
		}()
	}

	wg.Wait()
}
