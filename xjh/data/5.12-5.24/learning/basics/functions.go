package basics

import (
	"fmt"
	"sort"
	"strings"
)

// 函数进阶示例

// 递归函数示例
func RecursionExamples() {
	// 阶乘递归
	fmt.Println("5的阶乘:", factorial(5))

	// 斐波那契递归
	fmt.Println("斐波那契数列第10个数:", fibonacci(10))

	// 汉诺塔递归
	fmt.Println("汉诺塔移动步骤:")
	hanoi(3, "A", "B", "C")
}

// 阶乘函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 汉诺塔递归算法
func hanoi(n int, source, auxiliary, target string) {
	if n == 1 {
		fmt.Printf("将盘子 1 从 %s 移动到 %s\n", source, target)
		return
	}
	hanoi(n-1, source, target, auxiliary)
	fmt.Printf("将盘子 %d 从 %s 移动到 %s\n", n, source, target)
	hanoi(n-1, auxiliary, source, target)
}

// 高阶函数示例
func HigherOrderFunctionExamples() {
	// 函数作为参数
	numbers := []int{5, 2, 8, 1, 9}

	fmt.Println("原始数组:", numbers)

	fmt.Println("过滤后（只保留偶数）:", filter(numbers, isEven))
	fmt.Println("过滤后（只保留奇数）:", filter(numbers, isOdd))
	fmt.Println("过滤后（只保留大于5的数）:", filter(numbers, greaterThanFive))

	fmt.Println("映射后（每个数平方）:", mapFunc(numbers, square))
	fmt.Println("映射后（每个数乘以10）:", mapFunc(numbers, multiplyByTen))

	fmt.Println("求和:", reduce(numbers, 0, sum))
	fmt.Println("求积:", reduce(numbers, 1, product))

	// 闭包创建函数
	greaterThan3 := makeGreaterThan(3)
	greaterThan7 := makeGreaterThan(7)

	fmt.Println("大于3的数:", filter(numbers, greaterThan3))
	fmt.Println("大于7的数:", filter(numbers, greaterThan7))
}

// 过滤函数
func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, v := range numbers {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// 映射函数
func mapFunc(numbers []int, f func(int) int) []int {
	result := make([]int, len(numbers))
	for i, v := range numbers {
		result[i] = f(v)
	}
	return result
}

// 规约函数
func reduce(numbers []int, initialValue int, f func(int, int) int) int {
	result := initialValue
	for _, v := range numbers {
		result = f(result, v)
	}
	return result
}

// 判断偶数的函数
func isEven(n int) bool {
	return n%2 == 0
}

// 判断奇数的函数
func isOdd(n int) bool {
	return n%2 != 0
}

// 判断大于5的函数
func greaterThanFive(n int) bool {
	return n > 5
}

// 平方函数
func square(n int) int {
	return n * n
}

// 乘以10的函数
func multiplyByTen(n int) int {
	return n * 10
}

// 求和函数
func sum(acc, n int) int {
	return acc + n
}

// 求积函数
func product(acc, n int) int {
	return acc * n
}

// 闭包工厂函数
func makeGreaterThan(threshold int) func(int) bool {
	return func(n int) bool {
		return n > threshold
	}
}

// 函数式编程示例
func FunctionalProgrammingExamples() {
	// 使用函数式方法处理数据
	people := []struct {
		name string
		age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
		{"Dave", 35},
		{"Eve", 28},
	}

	// 使用sort包的函数式排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println("按年龄排序:")
	for _, p := range people {
		fmt.Printf("%s: %d岁\n", p.name, p.age)
	}

	// 使用自定义的函数式方法
	names := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}

	// 1. 将所有名字转为小写
	lowercaseNames := transformStrings(names, strings.ToLower)
	fmt.Println("小写转换:", lowercaseNames)

	// 2. 筛选出长度大于3的名字
	longNames := filterStrings(names, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println("长度大于3的名字:", longNames)

	// 3. 链式操作：先筛选，再转换
	result := transformStrings(
		filterStrings(names, func(s string) bool {
			return len(s) > 3
		}),
		func(s string) string {
			return s + "!"
		},
	)
	fmt.Println("链式操作结果:", result)

	// TODO: 尝试更多函数式编程示例
}

// 字符串转换函数
func transformStrings(strs []string, f func(string) string) []string {
	result := make([]string, len(strs))
	for i, s := range strs {
		result[i] = f(s)
	}
	return result
}

// 字符串过滤函数
func filterStrings(strs []string, f func(string) bool) []string {
	var result []string
	for _, s := range strs {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

// 方法与接收器示例
func MethodsAndReceiversExamples() {
	// 创建矩形实例
	r := Rectangle{width: 5, height: 3}
	fmt.Printf("矩形: 宽度=%d, 高度=%d\n", r.width, r.height)
	fmt.Printf("面积: %.2f\n", r.Area())
	fmt.Printf("周长: %d\n", r.Perimeter())

	// 修改矩形属性
	r.Scale(2)
	fmt.Printf("缩放后的矩形: 宽度=%d, 高度=%d\n", r.width, r.height)
	fmt.Printf("新面积: %.2f\n", r.Area())

	// 使用值接收器和指针接收器
	c1 := Counter{value: 10}
	c2 := Counter{value: 10}

	c1.IncrementByValue()     // 值接收器
	c2.IncrementByReference() // 指针接收器

	fmt.Printf("值接收器增加后: %d\n", c1.value)  // 应该仍为10
	fmt.Printf("指针接收器增加后: %d\n", c2.value) // 应该为11

	// 多次调用
	ptr := &c1
	for i := 0; i < 3; i++ {
		ptr.IncrementByReference()
	}
	fmt.Printf("调用3次指针方法后: %d\n", c1.value)

	// TODO: 尝试更多方法和接收器示例
}

// 矩形结构体
type Rectangle struct {
	width  int
	height int
}

// 计算面积的方法（值接收器）
func (r Rectangle) Area() float64 {
	return float64(r.width * r.height)
}

// 计算周长的方法（值接收器）
func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

// 缩放的方法（指针接收器）
func (r *Rectangle) Scale(factor int) {
	r.width *= factor
	r.height *= factor
}

// 计数器结构体
type Counter struct {
	value int
}

// 值接收器的方法（不会改变原始值）
func (c Counter) IncrementByValue() {
	c.value++
}

// 指针接收器的方法（会改变原始值）
func (c *Counter) IncrementByReference() {
	c.value++
}

// 接口和多态示例
func InterfacesAndPolymorphismExamples() {
	// 创建不同形状
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
		Triangle{base: 3, height: 4},
	}

	// 多态调用Area()方法
	fmt.Println("各形状的面积:")
	for i, shape := range shapes {
		fmt.Printf("形状%d: %T, 面积: %.2f\n", i+1, shape, shape.Area())
	}

	// 类型断言
	fmt.Println("\n使用类型断言获取具体类型:")
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("圆形 - 半径: %d, 面积: %.2f\n", s.radius, s.Area())
		case Rectangle:
			fmt.Printf("矩形 - 宽度: %d, 高度: %d, 面积: %.2f\n",
				s.width, s.height, s.Area())
		case Triangle:
			fmt.Printf("三角形 - 底: %d, 高: %d, 面积: %.2f\n",
				s.base, s.height, s.Area())
		default:
			fmt.Printf("未知形状: %T\n", s)
		}
	}

	// 空接口作为任意类型
	var anything interface{}
	anything = 42
	fmt.Println("\n空接口存储整数:", anything)

	anything = "Hello, Go!"
	fmt.Println("空接口存储字符串:", anything)

	anything = []float64{1.1, 2.2, 3.3}
	fmt.Println("空接口存储切片:", anything)

	// 类型断言
	value, ok := anything.([]float64)
	if ok {
		fmt.Println("类型断言成功, 值:", value)
	}

	// TODO: 尝试更多接口和多态示例
}

// 形状接口
type Shape interface {
	Area() float64
}

// 圆形结构体
type Circle struct {
	radius int
}

// 圆形的面积计算方法
func (c Circle) Area() float64 {
	return 3.14 * float64(c.radius) * float64(c.radius)
}

// 三角形结构体
type Triangle struct {
	base   int
	height int
}

// 三角形的面积计算方法
func (t Triangle) Area() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}
