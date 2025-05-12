package main

import (
	"flag"
	"fmt"
	"os"

	"learning/advanced"
	"learning/basics"
	// "learning/datastructs" // 等datastructs模块完成后再取消注释
)

func main() {
	// 定义命令行参数
	var (
		showAll         = flag.Bool("all", false, "运行所有示例")
		showBasics      = flag.Bool("basics", false, "运行基础语法示例")
		showAdvanced    = flag.Bool("advanced", false, "运行高级特性示例")
		showDataStructs = flag.Bool("datastructs", false, "运行数据结构示例")
	)

	// 解析命令行参数
	flag.Parse()

	// 如果没有提供参数，显示使用帮助
	if !*showAll && !*showBasics && !*showAdvanced && !*showDataStructs && flag.NArg() == 0 {
		fmt.Println("Go语言学习示例程序")
		fmt.Println("\n使用方法:")
		fmt.Println("  go run main.go [选项] [示例名称]")
		fmt.Println("\n选项:")
		flag.PrintDefaults()
		fmt.Println("\n示例名称:")
		fmt.Println("  variables    - 变量声明和使用")
		fmt.Println("  datatypes    - 数据类型")
		fmt.Println("  control      - 控制流")
		fmt.Println("  functions    - 函数")
		fmt.Println("  errors       - 错误处理")
		fmt.Println("  recursion    - 递归函数")
		fmt.Println("  higher       - 高阶函数")
		fmt.Println("  functional   - 函数式编程")
		fmt.Println("  methods      - 方法和接收器")
		fmt.Println("  interfaces   - 接口和多态")
		fmt.Println("  reflection   - 反射")
		fmt.Println("  advanced     - 高级特性")
		os.Exit(0)
	}

	// 根据参数运行相应的示例
	if *showAll || *showBasics {
		runBasicsExamples()
	}

	if *showAll || *showAdvanced {
		runAdvancedExamples()
	}

	if *showAll || *showDataStructs {
		runDataStructsExamples()
	}

	// 运行特定示例
	if flag.NArg() > 0 {
		runSpecificExample(flag.Arg(0))
	}
}

func runBasicsExamples() {
	fmt.Println("\n===== 基础语法示例 =====")

	fmt.Println("\n--- 变量示例 ---")
	basics.VariableExamples()

	fmt.Println("\n--- 数据类型示例 ---")
	basics.DataTypesExamples()

	fmt.Println("\n--- 控制流示例 ---")
	basics.ControlFlowExamples()

	fmt.Println("\n--- 函数示例 ---")
	basics.FunctionExamples()

	fmt.Println("\n--- 错误处理示例 ---")
	basics.ErrorHandlingExamples()
}

func runAdvancedExamples() {
	fmt.Println("\n===== 高级特性示例 =====")

	fmt.Println("\n--- 反射示例 ---")
	advanced.ReflectionDemo()

	fmt.Println("\n--- 高级特性集合 ---")
	advanced.AdvancedDemo()
}

func runDataStructsExamples() {
	fmt.Println("\n===== 数据结构示例 =====")
	// 等datastructs模块完成后再实现
	fmt.Println("数据结构示例尚未实现")
}

func runSpecificExample(name string) {
	switch name {
	case "variables":
		basics.VariableExamples()
	case "datatypes":
		basics.DataTypesExamples()
	case "control":
		basics.ControlFlowExamples()
	case "functions":
		basics.FunctionExamples()
	case "errors":
		basics.ErrorHandlingExamples()
	case "recursion":
		basics.RecursionExamples()
	case "higher":
		basics.HigherOrderFunctionExamples()
	case "functional":
		basics.FunctionalProgrammingExamples()
	case "methods":
		basics.MethodsAndReceiversExamples()
	case "interfaces":
		basics.InterfacesAndPolymorphismExamples()
	case "reflection":
		advanced.ReflectionDemo()
	case "advanced":
		advanced.AdvancedDemo()
	default:
		fmt.Printf("未知示例: %s\n", name)
		fmt.Println("请使用 'go run main.go' 查看可用示例列表")
	}
}
