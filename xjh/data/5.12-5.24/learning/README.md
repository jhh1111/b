# Go语言学习项目

这个项目旨在通过实践示例帮助你系统地学习Go语言的核心概念和高级特性。

## 目录结构

- **basics/** - Go语言基础知识
  - 基本语法和数据类型
  - 控制流和函数
  - 错误处理

- **advanced/** - Go语言高级特性
  - 并发编程与goroutine
  - 内存模型
  - 垃圾回收机制
  
- **datastructs/** - 数据结构实现与原理
  - 切片和映射底层原理
  - 切片操作技巧
  - 自定义数据结构
  
- **exercises/** - 练习任务
  - 每个主题的练习和解决方案
  - 阶段性项目

## 如何运行示例

项目提供了一个统一的命令行界面来运行各种示例。在项目根目录下运行以下命令：

```bash
go run main.go [选项] [示例名称]
```

### 命令行选项

- `-all` : 运行所有示例
- `-basics` : 运行所有基础语法示例
- `-advanced` : 运行所有高级特性示例
- `-datastructs` : 运行所有数据结构示例

### 示例名称列表

可以直接指定要运行的具体示例：

- `variables` - 变量声明和使用
- `datatypes` - 数据类型
- `control` - 控制流
- `functions` - 函数
- `errors` - 错误处理
- `recursion` - 递归函数
- `higher` - 高阶函数
- `functional` - 函数式编程
- `methods` - 方法和接收器
- `interfaces` - 接口和多态
- `reflection` - 反射
- `advanced` - 高级特性

### 示例运行命令

```bash
# 运行所有基础示例
go run main.go -basics

# 运行所有高级特性示例
go run main.go -advanced

# 运行特定示例
go run main.go functions
go run main.go interfaces
go run main.go reflection

# 运行所有示例
go run main.go -all
```

## 学习路径

### 第1阶段：Go语言基础 
- 掌握Go语言基础语法
- 理解基本数据类型和结构
- 熟悉控制流和函数定义
- **TODO**: 完成`basics/syntax.go`中的练习
- **TODO**: 实现`basics/functions.go`中的函数

### 第2阶段：Go数据结构 
- 深入理解切片和映射
- 掌握切片操作技巧
- 学习接口和结构体
- **TODO**: 完成`datastructs/slicetricks.go`中的函数实现
- **TODO**: 实现`datastructs/mapoperations.go`中的映射操作

### 第3阶段：并发编程 
- 学习goroutine和channel
- 理解同步原语(Mutex, WaitGroup等)
- 掌握context包使用
- **TODO**: 实现`advanced/concurrency.go`中的并发模式
- **TODO**: 完成`advanced/channels.go`中的channel操作

### 第4阶段：高级特性
- 学习反射和类型系统
- 了解内存模型和GC机制
- 掌握错误处理最佳实践
- **TODO**: 完成`advanced/reflection.go`中的反射练习
- **TODO**: 实现`advanced/memorymodel.go`中的内存优化

### 第5阶段：项目实践 
- 综合应用所学知识
- 实现爬虫项目核心组件
- 学习标准库使用
- **TODO**: 实现`exercises/crawler_components/fetcher.go`
- **TODO**: 完成`exercises/crawler_components/parser.go`

## 学习资源

- [Go语言官方教程](https://tour.golang.org/)
- [Go语言圣经](https://gopl.io/)
- [Go语言高级编程](https://github.com/chai2010/advanced-go-programming-book)

## 如何使用

1. 按照学习路径顺序学习各个模块
2. 完成每个文件中标记的TODO任务
3. 运行测试验证你的实现
4. 参考解决方案比较你的代码

## 测试和验证

每个练习都附带测试代码，可以通过以下命令运行测试：

```bash
cd 相应目录
go test
```

## 示例内容说明

### 基础语法 (basics)

- **变量** - 展示Go中不同的变量声明和初始化方式
- **数据类型** - 演示Go支持的基本数据类型和操作
- **控制流** - 包含if, for, switch, goto等控制结构
- **函数** - 普通函数、多返回值、命名返回值、可变参数函数等
- **错误处理** - 错误处理最佳实践和模式

### 高级特性 (advanced)

- **反射** - 展示如何使用反射检查类型、调用方法、创建对象等
- **内存模型** - 解释Go内存模型、逃逸分析、GC工作原理
- **并发编程** - 展示goroutine、channel、同步原语、并发模式等

### 数据结构 (datastructs)

- **切片技巧** - 展示Go切片的常用操作和高级技巧
- **映射操作** - 展示Go map的使用技巧和底层实现
