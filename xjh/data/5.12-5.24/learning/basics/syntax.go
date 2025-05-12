package basics

import (
	"fmt"
	"strconv"
)

// Go语言基础语法示例

// 变量声明和初始化
func VariableExamples() {
	// 1. 完整声明
	var a int = 10
	var b string = "hello"
	var c bool = true
	fmt.Println("完整声明:", a, b, c)

	// 2. 类型推断
	var d = 20
	var e = "world"
	var f = false
	fmt.Println("类型推断:", d, e, f)

	// 3. 短变量声明
	g := 30
	h := "go"
	i := true
	fmt.Println("短变量声明:", g, h, i)

	// 4. 多变量声明
	var j, k int = 40, 50
	var m, n = 60, "golang"
	o, p := 70, false
	fmt.Println("多变量声明:", j, k, m, n, o, p)

	// 5. 变量默认值
	var q int
	var r string
	var s bool
	fmt.Println("默认值:", q, r, s)

	// TODO: 尝试其他变量声明方式
}

// 基本数据类型
func DataTypesExamples() {
	// 1. 整数类型
	var a int = 10
	var b int8 = 127
	var c int16 = 32767
	var d int32 = 2147483647
	var e int64 = 9223372036854775807
	var f uint8 = 255
	fmt.Println("整数类型:", a, b, c, d, e, f)

	// 2. 浮点类型
	var g float32 = 3.14
	var h float64 = 3.141592653589793
	fmt.Println("浮点类型:", g, h)

	// 3. 复数类型
	var i complex64 = 1 + 2i
	var j complex128 = 1.1 + 2.2i
	fmt.Println("复数类型:", i, j)

	// 4. 布尔类型
	var k bool = true
	var l bool = false
	fmt.Println("布尔类型:", k, l)

	// 5. 字符串类型
	var m string = "你好，Go语言"
	fmt.Println("字符串类型:", m)

	// 6. 字节类型
	var n byte = 65 // ASCII码的'A'
	var o byte = 'A'
	fmt.Println("字节类型:", n, o, string(o))

	// 7. 符文类型
	var p rune = '中'
	fmt.Printf("符文类型: %d, %c\n", p, p)

	// TODO: 尝试不同数据类型的转换
}

// 控制流
func ControlFlowExamples() {
	// 1. if 条件语句
	a := 10
	if a > 0 {
		fmt.Println("a是正数")
	} else if a < 0 {
		fmt.Println("a是负数")
	} else {
		fmt.Println("a是零")
	}

	// 带初始化语句的if
	if b := 5; b > 0 {
		fmt.Println("b是正数:", b)
	}

	// 2. for 循环
	// 标准for循环
	for i := 0; i < 3; i++ {
		fmt.Println("标准for循环:", i)
	}

	// while风格for循环
	c := 0
	for c < 3 {
		fmt.Println("while风格for循环:", c)
		c++
	}

	// 无限循环
	d := 0
	for {
		fmt.Println("无限循环:", d)
		d++
		if d >= 3 {
			break
		}
	}

	// 遍历切片
	numbers := []int{1, 2, 3}
	for i, num := range numbers {
		fmt.Printf("下标: %d, 值: %d\n", i, num)
	}

	// 遍历映射
	colors := map[string]string{
		"red":   "红色",
		"green": "绿色",
		"blue":  "蓝色",
	}
	for k, v := range colors {
		fmt.Printf("键: %s, 值: %s\n", k, v)
	}

	// 3. switch语句
	day := "星期一"
	switch day {
	case "星期一":
		fmt.Println("周一快乐")
	case "星期五":
		fmt.Println("周五愉快")
	default:
		fmt.Println("普通的一天")
	}

	// 不带表达式的switch
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 4. goto语句
	i := 0
loop:
	fmt.Println("goto示例:", i)
	i++
	if i < 3 {
		goto loop
	}

	// TODO: 实现更复杂的控制流例子
}

// 函数示例
func FunctionExamples() {
	// 1. 基本函数调用
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// 2. 多返回值
	sum, diff := calculateSumAndDiff(10, 5)
	fmt.Println("10 + 5 =", sum, "10 - 5 =", diff)

	// 3. 命名返回值
	perimeter, area := calculateRectangle(5, 3)
	fmt.Println("矩形周长:", perimeter, "面积:", area)

	// 4. 可变参数
	total := sumNumbers(1, 2, 3, 4, 5)
	fmt.Println("1到5的和:", total)

	// 5. 匿名函数
	f := func(x, y int) int {
		return x * y
	}
	fmt.Println("3 x 4 =", f(3, 4))

	// 立即执行的匿名函数
	func(msg string) {
		fmt.Println("匿名函数输出:", msg)
	}("Hello Go!")

	// 6. 函数作为参数
	calculate(5, 3, add)
	calculate(5, 3, subtract)

	// 7. 闭包函数
	counter := createCounter()
	fmt.Println("计数器:", counter()) // 1
	fmt.Println("计数器:", counter()) // 2
	fmt.Println("计数器:", counter()) // 3

	// TODO: 尝试更多函数相关的示例
}

// 基本加法函数
func add(x, y int) int {
	return x + y
}

// 减法函数
func subtract(x, y int) int {
	return x - y
}

// 多返回值函数
func calculateSumAndDiff(x, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}

// 命名返回值函数
func calculateRectangle(length, width int) (perimeter int, area int) {
	perimeter = 2 * (length + width)
	area = length * width
	return // 隐式返回命名变量
}

// 可变参数函数
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 函数作为参数
func calculate(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Printf("计算结果: %d\n", result)
}

// 返回闭包函数
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 错误处理示例
func ErrorHandlingExamples() {
	// 1. 基本错误处理
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 尝试除以零
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	// 2. 字符串转换错误处理
	number, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println("转换结果:", number)
	}

	// 无效字符串转换
	number, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println("转换结果:", number)
	}

	// TODO: 尝试更多错误处理模式
}

// 可能返回错误的函数
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}
