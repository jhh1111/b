package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// SliceTricksDemo 展示切片的各种高级技巧和内部原理
func SliceTricksDemo() {
	fmt.Println("\n===== 切片高级技巧演示 =====")

	// TODO: 学生练习 - 实现下面的函数
	demoDeleteElement()
	demoPopElements()
	demoFilterElements()
	demoMapReduceFilter()
	demoSliceInternals()
}

// 从切片中删除元素
func demoDeleteElement() {
	fmt.Println("\n1. 从切片中删除元素:")

	// 示例数据
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", s)

	// 删除第三个元素 (索引2)
	i := 2
	fmt.Printf("删除索引 %d 的元素\n", i)

	// TODO: 实现删除元素方式1 - 保持顺序
	s1 := DeleteKeepOrder(s, i)
	fmt.Println("保持顺序删除后:", s1)

	// TODO: 实现删除元素方式2 - 不需要保持顺序(更高效)
	s2 := DeleteNoOrder(s, i)
	fmt.Println("不保持顺序删除后:", s2)
}

// 从切片弹出元素(类似栈或队列操作)
func demoPopElements() {
	fmt.Println("\n2. 栈和队列操作:")

	// 示例数据
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", s)

	// TODO: 实现从头部弹出(队列: 先进先出)
	val, newSlice := PopFromFront(s)
	fmt.Printf("从头部弹出: %d, 结果: %v\n", val, newSlice)

	// TODO: 实现从尾部弹出(栈: 后进先出)
	val, newSlice = PopFromBack(s)
	fmt.Printf("从尾部弹出: %d, 结果: %v\n", val, newSlice)

	// TODO: 实现从头部添加(入队)
	newSlice = PushToFront(s, 0)
	fmt.Printf("头部添加0: %v\n", newSlice)

	// TODO: 实现从尾部添加(入栈)
	newSlice = PushToBack(s, 6)
	fmt.Printf("尾部添加6: %v\n", newSlice)
}

// 过滤切片元素
func demoFilterElements() {
	fmt.Println("\n3. 过滤切片元素:")

	// 示例数据
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原始切片:", s)

	// TODO: 实现过滤偶数
	evens := FilterSlice(s, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println("偶数:", evens)

	// TODO: 实现过滤大于5的数
	greaterThan5 := FilterSlice(s, func(v int) bool {
		return v > 5
	})
	fmt.Println("大于5的数:", greaterThan5)

	// TODO: 创建(可选): 使用泛型实现通用的FilterFunc
	// 如果你的Go版本支持泛型(Go 1.18+)
	// 例如: FilterT[T any](s []T, f func(T) bool) []T
}

// Map-Reduce-Filter模式
func demoMapReduceFilter() {
	fmt.Println("\n4. Map-Reduce-Filter模式:")

	// 示例数据
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", s)

	// TODO: 实现Map函数 - 对每个元素应用变换
	squared := MapSlice(s, func(v int) int {
		return v * v
	})
	fmt.Println("Map(平方):", squared)

	// TODO: 实现Reduce函数 - 聚合所有元素
	sum := ReduceSlice(s, func(acc, v int) int {
		return acc + v
	}, 0)
	fmt.Println("Reduce(求和):", sum)

	// 连续操作示例: 先筛选偶数，再求平方，最后求和
	// TODO: 实现流式API或使用以上函数组合
}

// 展示切片的内部结构和内存布局
func demoSliceInternals() {
	fmt.Println("\n5. 切片内部结构:")

	// 创建一个底层数组
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 创建多个共享底层数组的切片
	s1 := arr[2:5]
	s2 := arr[4:7]
	s3 := s1[1:3]

	// 打印切片信息
	fmt.Println("arr:", arr)
	fmt.Printf("s1 = arr[2:5]: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = arr[4:7]: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3 = s1[1:3]: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// 展示内存共享
	fmt.Println("\n修改s3[0] = 99:")
	s3[0] = 99
	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	// TODO: 实现一个函数显示切片的内部结构信息
	fmt.Println("\n切片内部结构信息:")
	InspectSlice(s1, "s1")
	InspectSlice(s2, "s2")
	InspectSlice(s3, "s3")
}

// ===== 以下是需要学生实现的函数 =====

// DeleteKeepOrder 删除切片中的元素并保持顺序
// TODO: 学生需要实现这个函数
func DeleteKeepOrder(slice []int, index int) []int {
	// 保持顺序删除切片元素
	// 提示: 使用 append 和切片操作
	return nil
}

// DeleteNoOrder 删除切片中的元素但不保持顺序(更高效)
// TODO: 学生需要实现这个函数
func DeleteNoOrder(slice []int, index int) []int {
	// 不保持顺序删除切片元素(通常更高效)
	// 提示: 将最后一个元素复制到要删除的位置，然后截断切片
	return nil
}

// PopFromFront 从切片头部弹出元素
// TODO: 学生需要实现这个函数
func PopFromFront(slice []int) (int, []int) {
	// 从头部弹出元素并返回剩余切片
	// 提示: 记得处理空切片情况
	return 0, nil
}

// PopFromBack 从切片尾部弹出元素
// TODO: 学生需要实现这个函数
func PopFromBack(slice []int) (int, []int) {
	// 从尾部弹出元素并返回剩余切片
	// 提示: 记得处理空切片情况
	return 0, nil
}

// PushToFront 在切片头部添加元素
// TODO: 学生需要实现这个函数
func PushToFront(slice []int, value int) []int {
	// 在头部添加元素
	// 提示: 创建足够容量的新切片或使用append技巧
	return nil
}

// PushToBack 在切片尾部添加元素
// TODO: 学生需要实现这个函数
func PushToBack(slice []int, value int) []int {
	// 在尾部添加元素
	// 提示: 使用append
	return nil
}

// FilterSlice 根据条件过滤切片元素
// TODO: 学生需要实现这个函数
func FilterSlice(slice []int, keep func(int) bool) []int {
	// 返回符合条件的元素切片
	// 提示: 创建一个新切片，将符合条件的元素添加进去
	return nil
}

// MapSlice 对切片中的每个元素应用变换函数
// TODO: 学生需要实现这个函数
func MapSlice(slice []int, transform func(int) int) []int {
	// 对每个元素应用变换并返回新切片
	// 提示: 创建相同长度的新切片，填入变换后的值
	return nil
}

// ReduceSlice 将切片元素聚合为单一结果
// TODO: 学生需要实现这个函数
func ReduceSlice(slice []int, aggregate func(acc, val int) int, initial int) int {
	// 使用聚合函数将所有元素合并为一个结果
	// 提示: 从初始值开始，依次应用聚合函数
	return 0
}

// InspectSlice 检查切片的内部结构并打印详细信息
// TODO: 学生需要实现这个函数
func InspectSlice(slice []int, name string) {
	// 使用反射和unsafe包检查切片的内部结构
	// 打印切片的数据指针、长度、容量等信息
	// 提示: 需要使用unsafe.Pointer和reflect包
}

// ===== 以下是一些帮助函数的示例实现 =====

// SliceHeader 代表切片的运行时结构
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// GetSliceHeader 返回切片的头部信息
func GetSliceHeader(slice interface{}) SliceHeader {
	sh := SliceHeader{}

	// 使用reflect获取切片的值
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("不是切片类型")
	}

	// 使用unsafe.Pointer获取内部结构
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	// 填充SliceHeader
	sh.Data = unsafe.Pointer(header.Data)
	sh.Len = header.Len
	sh.Cap = header.Cap

	return sh
}

// 扩展学习:
// 1. 实现泛型版本(Go 1.18+)的切片操作
// 2. 探究切片扩容的内部机制(可以写一个测试函数)
// 3. 实现更高效的切片操作，考虑内存占用和GC压力
// 4. 实现线程安全的切片操作
// 5. 比较不同切片操作的性能
