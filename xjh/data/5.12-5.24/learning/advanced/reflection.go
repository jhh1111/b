package advanced

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// ReflectionDemo 演示Go语言的反射机制和类型系统
func ReflectionDemo() {
	fmt.Println("\n===== 反射(Reflection)和类型系统演示 =====")

	// 1. reflect.Type和reflect.Value
	fmt.Println("\n1. reflect.Type和reflect.Value基础:")

	// 使用简单的变量演示
	i := 42
	s := "hello, reflection"
	f := 3.14159

	// 获取类型
	fmt.Printf("i 的类型: %v\n", reflect.TypeOf(i))
	fmt.Printf("s 的类型: %v\n", reflect.TypeOf(s))
	fmt.Printf("f 的类型: %v\n", reflect.TypeOf(f))

	// 获取值
	fmt.Printf("i 的值: %v\n", reflect.ValueOf(i))
	fmt.Printf("s 的值: %v\n", reflect.ValueOf(s))
	fmt.Printf("f 的值: %v\n", reflect.ValueOf(f))

	// 2. 使用反射修改值
	fmt.Println("\n2. 使用反射修改值:")
	x := 100
	fmt.Printf("原始 x = %d\n", x)

	// 获取x的反射值
	v := reflect.ValueOf(&x).Elem()

	// 检查是否可设置
	fmt.Printf("v 是否可设置: %v\n", v.CanSet())

	// 修改值
	v.SetInt(200)
	fmt.Printf("修改后 x = %d\n", x)

	// 3. 反射结构体
	fmt.Println("\n3. 反射结构体:")
	p := Person{
		Name: "Alice",
		Age:  30,
		Job:  "Developer",
	}
	examineStruct(p)

	// 4. 反射调用方法
	fmt.Println("\n4. 反射调用方法:")
	reflectMethod(p)

	// 5. 接口底层实现
	fmt.Println("\n5. 接口底层实现:")
	examineInterface()

	// 6. 类型断言
	fmt.Println("\n6. 类型断言:")
	typeAssertionDemo()

	// 7. 反射创建新的值
	fmt.Println("\n7. 反射创建新的值:")
	createWithReflection()

	// TODO: 实现一个通用的JSON序列化器
	fmt.Println("\nTODO 练习: 实现一个基于反射的JSON序列化器")
	// 函数签名: func ToJSON(v interface{}) (string, error)
	// 要求支持基本类型、结构体、切片和映射

	// TODO: 实现一个通用的深拷贝函数
	fmt.Println("\nTODO 练习: 实现一个基于反射的深拷贝函数")
	// 函数签名: func DeepCopy(src interface{}) interface{}
}

// Person 用于反射演示的结构体
type Person struct {
	Name string
	Age  int
	Job  string `json:"occupation" description:"职业信息"`
}

// SayHello 是Person的方法
func (p Person) SayHello() string {
	return fmt.Sprintf("你好，我是%s，今年%d岁，我是一名%s", p.Name, p.Age, p.Job)
}

// UpdateAge 是Person的指针接收器方法
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// examineStruct 使用反射检查结构体
func examineStruct(obj interface{}) {
	// 获取类型信息
	t := reflect.TypeOf(obj)
	fmt.Printf("类型名: %s\n", t.Name())
	fmt.Printf("类别: %s\n", t.Kind())

	// 获取值信息
	v := reflect.ValueOf(obj)

	// 遍历字段
	fmt.Println("结构体字段:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("  %s: %v (类型: %v)\n", field.Name, value, field.Type)

		// 获取标签
		if tag := field.Tag; tag != "" {
			fmt.Printf("    标签: %s\n", tag)

			// 解析特定标签
			if jsonTag := tag.Get("json"); jsonTag != "" {
				fmt.Printf("    JSON标签: %s\n", jsonTag)
			}

			if descTag := tag.Get("description"); descTag != "" {
				fmt.Printf("    描述: %s\n", descTag)
			}
		}
	}

	// 遍历方法
	fmt.Println("结构体方法:")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("  %s: %v\n", method.Name, method.Type)
	}
}

// reflectMethod 使用反射调用方法
func reflectMethod(obj interface{}) {
	v := reflect.ValueOf(obj)
	m := v.MethodByName("SayHello")

	if m.IsValid() {
		// 调用无参数方法
		result := m.Call(nil)
		fmt.Printf("调用SayHello的结果: %v\n", result[0])
	}

	// 创建指针值，以调用指针接收器方法
	p := reflect.New(reflect.TypeOf(obj))
	p.Elem().Set(v) // 初始化

	// 获取UpdateAge方法
	updateMethod := p.MethodByName("UpdateAge")
	if updateMethod.IsValid() {
		// 调用有参数方法
		args := []reflect.Value{reflect.ValueOf(35)}
		updateMethod.Call(args)

		// 查看修改结果
		fmt.Printf("修改后年龄: %v\n", p.Elem().FieldByName("Age"))
	}
}

// 定义一个接口类型
type Speaker interface {
	Speak() string
}

// Dog 实现Speaker接口
type Dog struct {
	Name string
}

// Speak Dog实现Speaker接口
func (d Dog) Speak() string {
	return fmt.Sprintf("%s说：汪汪！", d.Name)
}

// Cat 实现Speaker接口
type Cat struct {
	Name string
}

// Speak Cat实现Speaker接口
func (c Cat) Speak() string {
	return fmt.Sprintf("%s说：喵喵～", c.Name)
}

// 接口的内部表示（模拟）
type iface struct {
	tab  *itab          // 类型信息
	data unsafe.Pointer // 数据指针
}

type itab struct {
	inter *interfacetype // 接口类型
	_type *_type         // 动态类型
}

type interfacetype struct {
	// ... 接口类型字段
}

type _type struct {
	// ... 类型信息字段
}

// examineInterface 探索接口的内部实现
func examineInterface() {
	var s Speaker

	fmt.Printf("空接口: %v, 类型: %T\n", s, s)

	// 赋值Dog
	s = Dog{Name: "旺财"}
	fmt.Printf("Dog接口值: %v, 类型: %T\n", s, s)
	fmt.Println(s.Speak())

	// 赋值Cat
	s = Cat{Name: "咪咪"}
	fmt.Printf("Cat接口值: %v, 类型: %T\n", s, s)
	fmt.Println(s.Speak())

	// 使用反射查看接口内部
	v := reflect.ValueOf(s)
	fmt.Printf("接口反射类型: %v, 种类: %v\n", v.Type(), v.Kind())

	// 解释接口的两个部分：类型表(type table)和数据指针(data pointer)
	fmt.Println("接口由两部分组成:")
	fmt.Println("1. 类型表(Type Table): 包含接口类型信息和动态类型的方法集")
	fmt.Println("2. 数据指针(Data Pointer): 指向实际数据的指针")

	// 显示内存布局
	fmt.Println("\n接口的内存布局示意图:")
	fmt.Println("+-----------------+")
	fmt.Println("| 类型表指针(tab) |")
	fmt.Println("+-----------------+")
	fmt.Println("| 数据指针(data)  |")
	fmt.Println("+-----------------+")
}

// typeAssertionDemo 演示类型断言
func typeAssertionDemo() {
	// 创建一个空接口变量
	var i interface{}

	// 断言基本类型
	i = 42
	if val, ok := i.(int); ok {
		fmt.Printf("i 是整数: %d\n", val)
	}

	// 不成功的断言
	if val, ok := i.(string); ok {
		fmt.Printf("i 是字符串: %s\n", val)
	} else {
		fmt.Println("i 不是字符串")
	}

	// 使用空接口存储不同类型
	values := []interface{}{
		42,
		"hello",
		true,
		3.14,
		[]int{1, 2, 3},
		map[string]int{"one": 1, "two": 2},
		Dog{Name: "小黑"},
	}

	// 使用switch进行类型断言
	fmt.Println("\n使用type switch进行类型断言:")
	for i, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("值 %d 是整数: %d\n", i, val)
		case string:
			fmt.Printf("值 %d 是字符串: %s\n", i, val)
		case bool:
			fmt.Printf("值 %d 是布尔值: %v\n", i, val)
		case float64:
			fmt.Printf("值 %d 是浮点数: %f\n", i, val)
		case []int:
			fmt.Printf("值 %d 是整数切片: %v\n", i, val)
		case map[string]int:
			fmt.Printf("值 %d 是映射: %v\n", i, val)
		case Dog:
			fmt.Printf("值 %d 是Dog: %s\n", i, val.Speak())
		default:
			fmt.Printf("值 %d 是未知类型: %T\n", i, val)
		}
	}

	// 接口嵌套的类型断言
	var s Speaker = Dog{Name: "旺财"}

	// 从接口类型断言为具体类型
	if dog, ok := s.(Dog); ok {
		fmt.Printf("s 是Dog: %s\n", dog.Name)
	}

	// 从接口类型断言为另一个接口类型
	if _, ok := s.(fmt.Stringer); ok {
		fmt.Println("s 同样实现了Stringer接口")
	} else {
		fmt.Println("s 没有实现Stringer接口")
	}
}

// createWithReflection 使用反射创建新的值
func createWithReflection() {
	// 1. 创建基本类型
	intType := reflect.TypeOf(0)
	intValue := reflect.New(intType).Elem()
	intValue.SetInt(42)
	fmt.Printf("创建整数: %v (%v)\n", intValue.Interface(), intValue.Type())

	// 2. 创建数组
	arrayType := reflect.ArrayOf(3, reflect.TypeOf(0))
	arrayValue := reflect.New(arrayType).Elem()
	for i := 0; i < 3; i++ {
		arrayValue.Index(i).SetInt(int64(i + 1))
	}
	fmt.Printf("创建数组: %v (%v)\n", arrayValue.Interface(), arrayValue.Type())

	// 3. 创建切片
	sliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf("")), 0, 3)
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("a"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("b"))
	sliceValue = reflect.Append(sliceValue, reflect.ValueOf("c"))
	fmt.Printf("创建切片: %v (%v)\n", sliceValue.Interface(), sliceValue.Type())

	// 4. 创建映射
	mapValue := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0)))
	mapValue.SetMapIndex(reflect.ValueOf("one"), reflect.ValueOf(1))
	mapValue.SetMapIndex(reflect.ValueOf("two"), reflect.ValueOf(2))
	fmt.Printf("创建映射: %v (%v)\n", mapValue.Interface(), mapValue.Type())

	// 5. 创建结构体
	structType := reflect.StructOf([]reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
		},
	})

	structValue := reflect.New(structType).Elem()
	structValue.Field(0).SetString("Dynamic Person")
	structValue.Field(1).SetInt(25)
	fmt.Printf("创建结构体: %v (%v)\n", structValue.Interface(), structValue.Type())
}

// SimpleSerializer 简单的对象序列化器，将其转换为字符串表示
func SimpleSerializer(v interface{}) string {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf("\"%s\"", value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", value.Bool())
	case reflect.Slice, reflect.Array:
		var items []string
		for i := 0; i < value.Len(); i++ {
			items = append(items, SimpleSerializer(value.Index(i).Interface()))
		}
		return "[" + strings.Join(items, ", ") + "]"
	case reflect.Map:
		var pairs []string
		iter := value.MapRange()
		for iter.Next() {
			k := SimpleSerializer(iter.Key().Interface())
			v := SimpleSerializer(iter.Value().Interface())
			pairs = append(pairs, fmt.Sprintf("%s: %s", k, v))
		}
		return "{" + strings.Join(pairs, ", ") + "}"
	case reflect.Struct:
		var fields []string
		t := value.Type()
		for i := 0; i < value.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := SimpleSerializer(value.Field(i).Interface())
			fields = append(fields, fmt.Sprintf("%s: %s", fieldName, fieldValue))
		}
		return t.Name() + " {" + strings.Join(fields, ", ") + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// TODO: 实现一个完整的ToJSON函数
// func ToJSON(v interface{}) (string, error) {
//     // 实现完整的序列化逻辑
// }

// TODO: 实现DeepCopy函数
// func DeepCopy(src interface{}) interface{} {
//     // 实现深拷贝逻辑
// }
