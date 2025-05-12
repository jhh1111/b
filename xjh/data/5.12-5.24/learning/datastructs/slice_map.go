package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// SliceDemo 演示Go语言切片的各种操作和底层原理
// 切片是Go中最常用的数据结构之一，它由三部分组成：
// 1. 指向底层数组的指针
// 2. 切片的长度
// 3. 切片的容量
func SliceDemo() {
	fmt.Println("\n===== 切片(Slice)演示 =====")

	// 创建切片的几种方式
	fmt.Println("\n1. 创建切片的方式:")

	// 1. 使用字面量直接创建
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1: %v, 长度: %d, 容量: %d\n", s1, len(s1), cap(s1))

	// 2. 使用make创建指定长度和容量的切片
	s2 := make([]int, 3, 5) // 长度为3，容量为5
	fmt.Printf("s2: %v, 长度: %d, 容量: %d\n", s2, len(s2), cap(s2))

	// 3. 从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:4] // 包含索引1到3的元素
	fmt.Printf("s3: %v, 长度: %d, 容量: %d\n", s3, len(s3), cap(s3))

	// 切片的底层结构展示
	fmt.Println("\n2. 切片的底层结构:")
	// 使用unsafe.Pointer和reflect查看切片的内部结构
	showSliceInternals(s1)

	// 切片共享底层数组
	fmt.Println("\n3. 切片共享底层数组:")
	s4 := s1[1:3]
	fmt.Printf("s1: %v, s4: %v\n", s1, s4)
	s4[0] = 100 // 修改s4的元素
	fmt.Printf("修改后 s1: %v, s4: %v (修改会影响原切片，因为共享底层数组)\n", s1, s4)

	// 切片扩容机制
	fmt.Println("\n4. 切片扩容机制:")
	s5 := make([]int, 0, 0)
	showCapGrowth(s5)

	// 切片的常用操作
	fmt.Println("\n5. 常用切片操作:")

	// 添加元素
	s6 := []int{1, 2, 3}
	fmt.Printf("原始切片: %v\n", s6)

	s6 = append(s6, 4) // 添加单个元素
	fmt.Printf("添加单个元素后: %v\n", s6)

	s6 = append(s6, 5, 6, 7) // 添加多个元素
	fmt.Printf("添加多个元素后: %v\n", s6)

	// 删除元素
	fmt.Printf("删除索引2的元素: %v\n", removeAt(s6, 2))

	// 插入元素
	fmt.Printf("在索引2插入元素99: %v\n", insertAt(s6, 2, 99))

	// 切片复制
	dst := make([]int, len(s6))
	copied := copy(dst, s6)
	fmt.Printf("复制切片: 复制了%d个元素, 结果: %v\n", copied, dst)

	// 切片作为函数参数(传引用行为)
	fmt.Println("\n6. 切片作为函数参数:")
	s7 := []int{1, 2, 3}
	fmt.Printf("调用前: %v\n", s7)
	modifySlice(s7)
	fmt.Printf("调用后: %v (切片作为参数传递时是引用传递行为)\n", s7)

	// TODO: 实现切片元素过滤函数
	// FilterSlice接受一个切片和一个过滤函数，返回一个新切片，包含所有满足条件的元素
	fmt.Println("\nTODO 练习: 实现切片元素过滤函数")
	// 要求实现filterSlice函数，使下面的代码能够正确工作：
	// filtered := filterSlice([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	// fmt.Printf("过滤后的切片(保留偶数): %v\n", filtered)

	// TODO: 实现自定义的切片排序
	// 要求使用sort.Interface接口实现自定义排序
	fmt.Println("\nTODO 练习: 实现自定义切片排序")
}

// MapDemo 演示Go语言映射(Map)的原理和操作
// Map是Go中的哈希表实现，用于存储键值对
// Map的内部结构包含桶(bucket)，每个桶可以存储多个键值对
func MapDemo() {
	fmt.Println("\n===== 映射(Map)演示 =====")

	// 1. 创建Map的几种方式
	fmt.Println("\n1. 创建Map的方式:")

	// 使用字面量创建
	m1 := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
	}
	fmt.Printf("m1: %v\n", m1)

	// 使用make创建
	m2 := make(map[string]int) // 创建空映射
	m2["apple"] = 5
	m2["banana"] = 8
	fmt.Printf("m2: %v\n", m2)

	// 映射的基本操作
	fmt.Println("\n2. Map基本操作:")

	// 添加/更新元素
	m1["grape"] = 10
	fmt.Printf("添加元素后: %v\n", m1)

	// 获取元素
	value, exists := m1["apple"]
	fmt.Printf("m1[\"apple\"] = %d, 存在: %t\n", value, exists)

	value, exists = m1["pear"]
	fmt.Printf("m1[\"pear\"] = %d, 存在: %t\n", value, exists)

	// 删除元素
	delete(m1, "orange")
	fmt.Printf("删除orange后: %v\n", m1)

	// 遍历映射
	fmt.Println("\n遍历映射:")
	for key, value := range m1 {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// Map的底层实现原理
	fmt.Println("\n3. Map的底层实现原理:")
	fmt.Println("- Go的map是哈希表实现的")
	fmt.Println("- 使用桶(bucket)存储键值对")
	fmt.Println("- 键的哈希值决定了它存储在哪个桶中")
	fmt.Println("- 发生哈希冲突时使用链表法解决")

	// Map的内存占用
	showMapMemory()

	// Map并发安全问题
	fmt.Println("\n4. Map的并发安全问题:")
	fmt.Println("- Map不是并发安全的，同时读写会导致panic")
	fmt.Println("- 并发场景应使用sync.Map或加锁保护")

	// 并发安全的Map示例 - 使用互斥锁实现
	safeMap := NewSafeMap()
	safeMap.Set("key1", "value1")
	value1, _ := safeMap.Get("key1")
	fmt.Printf("SafeMap: key1 = %v\n", value1)

	// TODO: 实现线程安全的缓存系统
	fmt.Println("\nTODO 练习: 实现一个线程安全的缓存系统")
	// 要求:
	// 1. 支持过期时间
	// 2. 支持最大容量限制
	// 3. 线程安全
	// 4. 实现LRU淘汰策略

	// TODO: 实现Map的深拷贝函数
	fmt.Println("\nTODO 练习: 实现Map的深拷贝函数")
	// 要求实现函数: func DeepCopyMap(m map[string]interface{}) map[string]interface{}
}

// 以下是工具函数的实现

// 使用unsafe.Pointer查看切片的内部结构
func showSliceInternals(s []int) {
	// 定义切片的运行时结构
	type sliceHeader struct {
		Data unsafe.Pointer
		Len  int
		Cap  int
	}

	// 获取切片的头部信息
	sh := (*sliceHeader)(unsafe.Pointer(&s))

	fmt.Printf("切片的内部结构:\n")
	fmt.Printf("  Data指针: %p\n", sh.Data)
	fmt.Printf("  Len(长度): %d\n", sh.Len)
	fmt.Printf("  Cap(容量): %d\n", sh.Cap)

	// 通过指针直接访问底层数组元素
	if len(s) > 0 {
		firstElemPtr := unsafe.Pointer(uintptr(sh.Data))
		firstElem := *(*int)(firstElemPtr)
		fmt.Printf("  第一个元素的值: %d\n", firstElem)
	}
}

// 演示切片扩容机制
func showCapGrowth(s []int) {
	fmt.Println("观察切片扩容过程:")

	prevCap := cap(s)
	fmt.Printf("  初始容量: %d\n", prevCap)

	// 不断添加元素观察容量变化
	for i := 0; i < 20; i++ {
		s = append(s, i)
		currCap := cap(s)

		if currCap != prevCap {
			growth := float64(currCap) / float64(prevCap)
			fmt.Printf("  从%d扩容到%d - 增长系数: %.2fx\n", prevCap, currCap, growth)
			prevCap = currCap
		}
	}

	fmt.Println("  小结: Go 1.18后通常按照1.25-2倍增长，具体规则较复杂")
}

// 从切片中移除指定索引的元素
func removeAt(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	// 通过重组切片来"删除"元素
	return append(slice[:index], slice[index+1:]...)
}

// 在指定索引位置插入元素
func insertAt(slice []int, index int, value int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	// 首先扩展切片长度
	slice = append(slice, 0)
	// 然后将index之后的元素向后移动
	copy(slice[index+1:], slice[index:])
	// 插入新值
	slice[index] = value
	return slice
}

// 修改切片内容的函数
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	// 注意: 虽然可以修改内容，但不能改变切片的长度
	// 这种改变在函数外是不可见的:
	s = append(s, 888)
}

// SafeMap 是一个线程安全的map实现
type SafeMap struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewSafeMap 创建一个新的线程安全映射
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]interface{}),
	}
}

// Get 安全地获取映射中的值
func (m *SafeMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[key]
	return val, ok
}

// Set 安全地设置映射中的值
func (m *SafeMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

// Delete 安全地删除映射中的键
func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
}

// Keys 返回映射中的所有键
func (m *SafeMap) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]string, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Len 返回映射的长度
func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

// 展示映射的内存占用
func showMapMemory() {
	fmt.Println("\nMap内存占用示例:")

	// 创建两个大小不同的map
	smallMap := make(map[int]int)
	largeMap := make(map[int]int)

	// 记录初始内存状态
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 向smallMap添加100个元素
	for i := 0; i < 100; i++ {
		smallMap[i] = i
	}

	// 记录中间内存状态
	runtime.ReadMemStats(&m2)
	smallMapSize := m2.Alloc - m1.Alloc

	// 向largeMap添加10000个元素
	for i := 0; i < 10000; i++ {
		largeMap[i] = i
	}

	// 记录最终内存状态
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	largeMapSize := m3.Alloc - m2.Alloc

	fmt.Printf("  smallMap (100个元素) 占用约: %d 字节\n", smallMapSize)
	fmt.Printf("  largeMap (10000个元素) 占用约: %d 字节\n", largeMapSize)
	fmt.Printf("  平均每个元素占用: %.2f 字节\n", float64(largeMapSize)/10000)
}

// 以下是未实现的TODO函数，需要学生自己完成

// TODO: 实现一个泛型的过滤器函数，接受一个切片和一个过滤函数，返回满足条件的元素
// func filterSlice[T any](slice []T, filterFn func(T) bool) []T {
//     // 实现过滤逻辑
// }

// TODO: 实现一个深拷贝Map的函数
// func DeepCopyMap(m map[string]interface{}) map[string]interface{} {
//     // 实现深拷贝逻辑
// }
