package advanced

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

// AdvancedDemo 展示Go语言的高级特性
func AdvancedDemo() {
	fmt.Println("\n===== Go高级特性演示 =====")

	// 1. 垃圾回收机制
	fmt.Println("\n1. 垃圾回收机制:")
	gcDemo()

	// 2. 内存逃逸分析
	fmt.Println("\n2. 内存逃逸分析:")
	escapeAnalysisDemo()

	// 3. Go内存模型
	fmt.Println("\n3. Go内存模型:")
	memoryModelDemo()

	// 4. sync包高级用法
	fmt.Println("\n4. sync包高级用法:")
	syncPackageDemo()

	// 5. atomic操作
	fmt.Println("\n5. 原子操作:")
	atomicOperationsDemo()

	// 6. channel高级模式
	fmt.Println("\n6. channel高级模式:")
	advancedChannelPatterns()

	// 7. GMP调度模型
	fmt.Println("\n7. GMP调度模型:")
	gmpSchedulerDemo()

	// TODO: 实现一个内存池
	fmt.Println("\nTODO 练习: 实现一个对象池来减少GC压力")

	// TODO: 实现高效的并发控制策略
	fmt.Println("\nTODO 练习: 实现一个带超时和取消功能的任务执行器")
}

// gcDemo 展示Go的垃圾回收机制
func gcDemo() {
	// 打印GC信息前先触发一次GC
	runtime.GC()
	printGCStats("初始状态")

	// 分配大量内存
	fmt.Println("\n分配大量内存:")
	{
		// 创建大量对象触发GC
		allocateObjects()

		// 打印GC状态
		printGCStats("分配对象后")

		// 手动触发GC
		fmt.Println("\n手动触发GC:")
		runtime.GC()
		printGCStats("手动GC后")
	}

	// 设置和获取GC参数
	fmt.Println("\nGC参数:")

	// 获取GOGC值
	gogc := debug.SetGCPercent(-1)
	fmt.Printf("原始GOGC值: %d%%\n", gogc)

	// 恢复GOGC值
	debug.SetGCPercent(gogc)

	// 设置GC最大暂停时间目标
	pauseTarget := 1000 * time.Microsecond // 1ms
	debug.SetGCPercent(500)                // 非常激进，为了演示
	debug.SetMaxStack(32 * 1024 * 1024)    // 32MB最大栈空间

	fmt.Printf("GC暂停时间目标: %v\n", pauseTarget)

	// GC的三色标记算法说明
	fmt.Println("\nGo GC的三色标记清除法:")
	fmt.Println("1. 白色: 潜在的垃圾，尚未被访问的对象")
	fmt.Println("2. 灰色: 已被访问但其引用尚未被扫描的对象")
	fmt.Println("3. 黑色: 已被访问且其所有引用也都已被扫描的对象")
	fmt.Println("4. 并发标记阶段会将对象从白色→灰色→黑色")
	fmt.Println("5. 清除阶段会回收所有白色对象")

	// 写屏障说明
	fmt.Println("\nGo GC的写屏障:")
	fmt.Println("- 写屏障是一种同步机制，确保在GC过程中对象引用的正确修改")
	fmt.Println("- Go使用混合写屏障(Hybrid Write Barrier)保证并发正确性")
	fmt.Println("- 写屏障在GC标记阶段开启，标记结束后关闭")
}

// allocateObjects 分配大量对象以触发GC
func allocateObjects() {
	// 分配多个大数组并立即丢弃引用
	for i := 0; i < 10; i++ {
		_ = make([]byte, 1024*1024) // 分配1MB
	}

	// 保留一些对象的引用防止被回收
	keeper = make([][]byte, 0)
	for i := 0; i < 5; i++ {
		keeper = append(keeper, make([]byte, 1024*1024)) // 保留5MB
	}
}

// keeper 用于防止对象被GC回收
var keeper [][]byte

// 打印GC统计信息
func printGCStats(label string) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Printf("=== %s ===\n", label)
	fmt.Printf("Alloc: %.2f MB\n", float64(stats.Alloc)/1024/1024)
	fmt.Printf("TotalAlloc: %.2f MB\n", float64(stats.TotalAlloc)/1024/1024)
	fmt.Printf("Sys: %.2f MB\n", float64(stats.Sys)/1024/1024)
	fmt.Printf("GC次数: %d\n", stats.NumGC)
	fmt.Printf("GC CPU占用: %.2f%%\n", float64(stats.GCCPUFraction)*100)
}

// escapeAnalysisDemo 演示Go的逃逸分析
// 要查看逃逸分析结果，可以使用: go build -gcflags="-m -l" advanced.go
func escapeAnalysisDemo() {
	// 1. 不逃逸的例子 - 本地变量不离开作用域
	fmt.Println("局部变量通常在栈上分配:")
	x := 42
	p := &x
	fmt.Printf("栈上变量x的地址: %p, 值: %d\n", p, *p)

	// 2. 逃逸的例子 - 返回局部变量的指针
	fmt.Println("\n返回局部变量指针会导致变量逃逸到堆:")
	result := createEscapedInt()
	fmt.Printf("堆上变量的地址: %p, 值: %d\n", result, *result)

	// 3. 闭包捕获变量导致逃逸
	fmt.Println("\n闭包捕获的变量会逃逸:")
	counter := createCounter()
	fmt.Printf("首次调用: %d\n", counter())
	fmt.Printf("再次调用: %d\n", counter())

	// 4. 接口类型赋值可能导致逃逸
	fmt.Println("\n分配给接口类型的值通常会逃逸:")
	var i interface{} = createValueForInterface()
	fmt.Printf("接口值的动态类型: %T\n", i)

	// 5. 大对象通常在堆上分配
	fmt.Println("\n大对象通常在堆上分配:")
	largeSlice := createLargeSlice()
	fmt.Printf("大切片的长度: %d\n", len(largeSlice))

	fmt.Println("\n逃逸分析的好处:")
	fmt.Println("- 减少GC压力：不逃逸的对象在栈上分配和回收，不需要GC")
	fmt.Println("- 提高性能：栈上分配比堆上分配更快")
	fmt.Println("- 减少内存碎片：栈内存使用更加紧凑")

	fmt.Println("\n使用go build -gcflags=\"-m -l\"查看逃逸分析详情")
}

// 返回局部变量的指针会导致变量逃逸
func createEscapedInt() *int {
	x := 100
	return &x // x会逃逸到堆上
}

// 闭包捕获变量导致逃逸
func createCounter() func() int {
	count := 0 // count会逃逸到堆上
	return func() int {
		count++
		return count
	}
}

// 返回值赋给接口类型导致逃逸
func createValueForInterface() int {
	return 200 // 返回值会逃逸，因为会被赋值给接口
}

// 大对象通常在堆上分配
func createLargeSlice() []int {
	return make([]int, 1000000) // 大对象通常在堆上分配
}

// memoryModelDemo 演示Go内存模型的重要概念
func memoryModelDemo() {
	fmt.Println("Go内存模型定义了goroutine间的内存可见性条件")

	// 演示不同步可能导致的问题
	fmt.Println("\n未同步的共享变量可能导致问题:")
	demonstrateRaceCondition()

	// 演示happen-before关系
	fmt.Println("\nHappen-before关系:")
	fmt.Println("- 同一goroutine中的语句按程序顺序happen-before")
	fmt.Println("- channel发送操作happen-before对应的接收操作完成")
	fmt.Println("- 互斥锁的解锁happen-before后续的加锁")
	fmt.Println("- 向无缓冲channel的发送happen-before对应的接收操作")

	demonstrateHappenBefore()

	// 内存重排序
	fmt.Println("\n内存重排序:")
	fmt.Println("- 编译器和CPU可能重排指令以优化性能")
	fmt.Println("- 代码中的顺序可能不是实际执行顺序")
	fmt.Println("- sync/atomic包提供的原子操作确保内存顺序")

	// 内存屏障
	fmt.Println("\n内存屏障:")
	fmt.Println("- sync.Mutex、atomic操作等会插入内存屏障")
	fmt.Println("- 内存屏障确保特定的内存操作顺序")
}

// 演示竞态条件
func demonstrateRaceCondition() {
	var counter int
	var wg sync.WaitGroup

	// 创建多个goroutine同时修改counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 没有同步机制保护，可能导致竞争
		}()
	}

	wg.Wait()
	fmt.Printf("未同步的计数器(期望1000): %d\n", counter)

	// 使用互斥锁解决
	var counterSafe int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counterSafe++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("同步的计数器(期望1000): %d\n", counterSafe)
}

// 演示happen-before关系
func demonstrateHappenBefore() {
	ch := make(chan int)

	go func() {
		x := 42
		// x的写入happen-before channel发送
		ch <- x
	}()

	// channel接收happen-before后续语句
	val := <-ch
	fmt.Printf("通过channel传递的值: %d\n", val)

	// 互斥锁的happen-before关系
	var mu sync.Mutex
	var sharedVar int

	go func() {
		mu.Lock()
		sharedVar = 100
		mu.Unlock() // 解锁happen-before后续的加锁
	}()

	time.Sleep(10 * time.Millisecond)

	mu.Lock() // 加锁happen-after前面的解锁
	fmt.Printf("通过互斥锁同步的值: %d\n", sharedVar)
	mu.Unlock()
}

// syncPackageDemo 展示sync包的高级用法
func syncPackageDemo() {
	// 1. sync.Once
	fmt.Println("sync.Once - 确保函数只执行一次:")
	demoSyncOnce()

	// 2. sync.Pool
	fmt.Println("\nsync.Pool - 对象池，减少内存分配:")
	demoSyncPool()

	// 3. sync.Map
	fmt.Println("\nsync.Map - 并发安全的map:")
	demoSyncMap()

	// 4. sync.Cond
	fmt.Println("\nsync.Cond - 条件变量:")
	demoSyncCond()

	// 5. 自旋锁vs互斥锁
	fmt.Println("\n自旋锁vs互斥锁:")
	fmt.Println("- 自旋锁: 冲突时CPU忙等待，适合锁竞争不激烈且持锁时间短")
	fmt.Println("- 互斥锁: 冲突时线程休眠，适合锁竞争激烈或持锁时间长")
	fmt.Println("- Go运行时会根据情况自动调整锁的行为")
}

// 演示sync.Once
func demoSyncOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	// 初始化函数
	initFunc := func() {
		fmt.Println("初始化操作执行(只会执行一次)")
	}

	// 多次调用
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutine %d 调用初始化函数\n", id)
			once.Do(initFunc) // 只有第一次调用会执行
		}(i)
	}

	wg.Wait()
}

// 演示sync.Pool
func demoSyncPool() {
	// 创建对象池
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("创建新的对象")
			return make([]byte, 1024) // 1KB缓冲区
		},
	}

	// 获取对象
	buf1 := pool.Get().([]byte)
	fmt.Printf("获取对象: %T, 长度: %d\n", buf1, len(buf1))

	// 修改对象并归还
	buf1[0] = 'A'
	pool.Put(buf1)

	// 再次获取(可能是之前的对象)
	buf2 := pool.Get().([]byte)
	fmt.Printf("再次获取: %T, 长度: %d, 首字节: %c\n", buf2, len(buf2), buf2[0])

	// 大量获取和归还
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 获取缓冲区
			buffer := pool.Get().([]byte)

			// 使用缓冲区
			// ...

			// 归还缓冲区
			pool.Put(buffer)
		}()
	}

	wg.Wait()
	fmt.Println("对象池操作完成")

	fmt.Println("\n对象池优点:")
	fmt.Println("- 减少GC压力: 复用对象而不是频繁创建")
	fmt.Println("- 提高性能: 避免频繁内存分配")
	fmt.Println("- 适用场景: 临时缓冲区、连接池、协程工作池等")
}

// 演示sync.Map
func demoSyncMap() {
	var m sync.Map

	// 存储键值对
	m.Store("name", "张三")
	m.Store("age", 30)
	m.Store("job", "程序员")

	// 获取值
	name, ok := m.Load("name")
	fmt.Printf("name: %v, 存在: %t\n", name, ok)

	// 不存在的键
	address, ok := m.Load("address")
	fmt.Printf("address: %v, 存在: %t\n", address, ok)

	// LoadOrStore - 如果键存在则返回值，否则存储并返回新值
	value, loaded := m.LoadOrStore("age", 40)
	fmt.Printf("age: %v, 已存在: %t\n", value, loaded)

	value, loaded = m.LoadOrStore("height", 175)
	fmt.Printf("height: %v, 已存在: %t\n", value, loaded)

	// 删除
	m.Delete("job")

	// 遍历
	fmt.Println("遍历sync.Map:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})
}

// 演示sync.Cond
func demoSyncCond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// 队列
	queue := make([]int, 0, 10)

	// 生产者
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(200 * time.Millisecond)

			mu.Lock()
			queue = append(queue, i)
			fmt.Printf("生产: %d\n", i)
			cond.Signal() // 通知一个等待的消费者
			mu.Unlock()
		}
	}()

	// 消费者
	for i := 0; i < 5; i++ {
		go func(id int) {
			mu.Lock()

			// 如果队列为空，等待
			for len(queue) == 0 {
				fmt.Printf("消费者%d等待...\n", id)
				cond.Wait() // 释放锁并等待通知
			}

			// 消费一个项目
			x := queue[0]
			queue = queue[1:]

			fmt.Printf("消费者%d获取: %d\n", id, x)

			mu.Unlock()
		}(i)
	}

	// 等待足够时间让所有操作完成
	time.Sleep(1500 * time.Millisecond)
}

// atomicOperationsDemo 演示原子操作
func atomicOperationsDemo() {
	var counter int64
	var wg sync.WaitGroup

	// 使用原子操作并发更新计数器
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("原子计数器最终值: %d\n", atomic.LoadInt64(&counter))

	// 比较并交换(CAS)
	fmt.Println("\n比较并交换(CAS)操作:")
	var value int32 = 100

	// 尝试CAS操作 - 匹配情况
	swapped := atomic.CompareAndSwapInt32(&value, 100, 200)
	fmt.Printf("第一次CAS: 值=%d, 交换成功=%t\n", atomic.LoadInt32(&value), swapped)

	// 尝试CAS操作 - 不匹配情况
	swapped = atomic.CompareAndSwapInt32(&value, 100, 300)
	fmt.Printf("第二次CAS: 值=%d, 交换成功=%t\n", atomic.LoadInt32(&value), swapped)

	// 原子操作类型
	fmt.Println("\n原子操作支持的类型:")
	fmt.Println("- 整数: int32, int64, uint32, uint64, uintptr")
	fmt.Println("- 指针: *T (任意类型T的指针)")

	// atomic.Value
	fmt.Println("\natomic.Value - 原子读写任意类型:")
	demoAtomicValue()
}

// 演示atomic.Value
func demoAtomicValue() {
	type Config struct {
		MaxConn int
		Timeout time.Duration
	}

	var config atomic.Value

	// 存储初始配置
	config.Store(Config{
		MaxConn: 10,
		Timeout: 1 * time.Second,
	})

	// 读取配置
	cfg := config.Load().(Config)
	fmt.Printf("初始配置: %+v\n", cfg)

	// 在另一个goroutine中更新配置
	go func() {
		time.Sleep(100 * time.Millisecond)

		// 原子地更新配置
		config.Store(Config{
			MaxConn: 20,
			Timeout: 2 * time.Second,
		})
	}()

	// 给更新足够的时间
	time.Sleep(200 * time.Millisecond)

	// 读取更新后的配置
	newCfg := config.Load().(Config)
	fmt.Printf("更新后配置: %+v\n", newCfg)
}

// advancedChannelPatterns 演示channel的高级使用模式
func advancedChannelPatterns() {
	// 1. Fan-out模式
	fmt.Println("Fan-out模式 - 多个goroutine从同一个channel接收数据:")
	demoFanOut()

	// 2. Fan-in模式
	fmt.Println("\nFan-in模式 - 多个channel合并到一个channel:")
	demoFanIn()

	// 3. 管道模式
	fmt.Println("\n管道模式 - channel连接的处理阶段:")
	demoPipeline()

	// 4. 超时和取消模式
	fmt.Println("\n超时和取消模式:")
	demoTimeoutAndCancel()

	// 5. 工作池模式
	fmt.Println("\n工作池模式:")
	demoWorkerPool()
}

// 演示Fan-out模式
func demoFanOut() {
	tasks := make(chan int, 10)

	// 发送任务
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// 启动3个worker
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for task := range tasks {
				fmt.Printf("worker %d 处理任务 %d\n", workerID, task)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("所有任务处理完成")
}

// 演示Fan-in模式
func demoFanIn() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	// 启动多个生产者
	go produce(c1, 1, 3)
	go produce(c2, 4, 6)
	go produce(c3, 7, 9)

	// 合并多个channel为一个
	merged := fanIn(c1, c2, c3)

	// 从合并后的channel接收数据
	for i := 0; i < 9; i++ {
		fmt.Printf("接收: %d\n", <-merged)
	}
}

// 生产者函数
func produce(ch chan<- int, start, end int) {
	for i := start; i <= end; i++ {
		ch <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

// 合并多个channel到一个channel
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// 为每个输入channel启动一个goroutine
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}

	// 所有输入处理完毕后关闭输出channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 演示管道模式
func demoPipeline() {
	// 创建管道的各个阶段
	nums := make(chan int, 5)
	squares := processStage(nums, func(x int) int { return x * x })
	final := processStage(squares, func(x int) int { return x + 1 })

	// 发送初始数据
	go func() {
		for i := 1; i <= 5; i++ {
			nums <- i
		}
		close(nums)
	}()

	// 接收处理后的数据
	fmt.Println("管道处理结果:")
	for result := range final {
		fmt.Printf("  %d\n", result)
	}
}

// 处理阶段
func processStage(in <-chan int, fn func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		for x := range in {
			out <- fn(x)
		}
		close(out)
	}()

	return out
}

// 演示超时和取消模式
func demoTimeoutAndCancel() {
	// 创建一个长时间运行的任务
	operation := func(cancel <-chan struct{}) <-chan string {
		result := make(chan string)

		go func() {
			defer close(result)

			// 模拟耗时操作
			for i := 0; i < 5; i++ {
				// 检查取消信号
				select {
				case <-cancel:
					fmt.Println("操作被取消")
					return
				default:
					// 继续执行
				}

				// 模拟工作
				time.Sleep(200 * time.Millisecond)
				fmt.Println("  操作进度:", i+1)
			}

			result <- "操作完成"
		}()

		return result
	}

	// 1. 超时示例
	fmt.Println("超时示例:")
	timeout := 500 * time.Millisecond
	cancellationCh := make(chan struct{})

	select {
	case res := <-operation(cancellationCh):
		fmt.Println("结果:", res)
	case <-time.After(timeout):
		close(cancellationCh) // 发送取消信号
		fmt.Println("操作超时")
	}

	// 2. 用户取消示例
	fmt.Println("\n用户取消示例:")
	cancellationCh = make(chan struct{})
	resultCh := operation(cancellationCh)

	// 模拟用户500ms后取消操作
	time.AfterFunc(500*time.Millisecond, func() {
		fmt.Println("用户取消操作")
		close(cancellationCh)
	})

	// 等待结果或取消
	if res, ok := <-resultCh; ok {
		fmt.Println("结果:", res)
	} else {
		fmt.Println("操作已取消")
	}

	// 使用context包更好地处理取消
	fmt.Println("\n使用context取消:")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("ctx取消原因:", ctx.Err())
	case <-time.After(1 * time.Second):
		fmt.Println("操作完成(不会执行到这里)")
	}
}

// 演示工作池模式
func demoWorkerPool() {
	// 任务和结果通道
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// 启动固定数量的worker
	numWorkers := 3
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// 发送任务
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	// 等待所有worker完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	fmt.Println("工作池结果:")
	for result := range results {
		fmt.Printf("  结果: %d\n", result)
	}
}

// 工作池的worker
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d 处理任务 %d\n", id, task)

		// 模拟处理任务
		time.Sleep(100 * time.Millisecond)

		// 发送结果
		results <- task * 2
	}
}

// gmpSchedulerDemo 展示GMP调度模型
func gmpSchedulerDemo() {
	fmt.Println("Go的GMP调度模型:")
	fmt.Println("- G: Goroutine，轻量级线程，由Go运行时管理")
	fmt.Println("- M: Machine，工作线程，关联一个操作系统线程")
	fmt.Println("- P: Processor，处理器，包含执行Go代码的资源")

	fmt.Println("\nGMP工作流程:")
	fmt.Println("1. P从本地队列获取G")
	fmt.Println("2. P将G分配给M执行")
	fmt.Println("3. G执行完毕后返回P的本地队列或全局队列")
	fmt.Println("4. 如果G阻塞，M会与P分离，P寻找其他M继续执行队列中的G")

	// 展示当前的goroutine数量
	fmt.Printf("\n当前goroutine数量: %d\n", runtime.NumGoroutine())

	// 展示可用的CPU核心数
	fmt.Printf("系统CPU核心数: %d\n", runtime.NumCPU())

	// 展示当前GOMAXPROCS设置
	gomaxprocs := runtime.GOMAXPROCS(0)
	fmt.Printf("当前GOMAXPROCS值: %d\n", gomaxprocs)

	// 创建许多goroutine，观察调度
	fmt.Println("\n创建100个goroutine观察调度:")
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 模拟CPU工作
			start := time.Now()
			for time.Since(start) < 10*time.Millisecond {
				// 空循环消耗CPU
			}
		}(i)
	}

	fmt.Printf("创建goroutine后，当前数量: %d\n", runtime.NumGoroutine())

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Printf("所有goroutine完成后，当前数量: %d\n", runtime.NumGoroutine())
}

// TODO: 实现一个限速器
// RateLimiter 接口定义
type RateLimiter interface {
	Allow() bool                    // 判断是否允许下一个请求
	Wait(ctx context.Context) error // 等待直到允许下一个请求或上下文取消
	Rate() float64                  // 返回限制速率
}

// TODO: 实现一个令牌桶算法的限速器
// func NewTokenBucketLimiter(rate float64, capacity int) RateLimiter {
//     // 实现令牌桶算法
// }

// TODO: 实现一个自定义内存池来减少GC压力
// ObjectPool 是一个通用的对象池
// type ObjectPool[T any] struct {
//     // 实现对象池
// }

// TODO: 实现一个对象的循环复用器
// func NewObjectPool[T any](newFn func() T, capacity int) *ObjectPool[T] {
//     // 实现对象池创建
// }
