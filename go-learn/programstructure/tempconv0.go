package main

import "fmt"

// 即使使用了相同的底层类型float64，它们不是相同的类型，不能使用算术表达式进行比较和合并
// 命名类型的底层类型决定了它的结构和表达方式，以及它支持的内部操作集合
type Celsius float64    // 摄氏温度类型
type Fahrenheit float64 // 华氏温度类型

const (
	AbsoluteZeroC Celsius = -273.15
	FreezlingC    Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// Celsius和Fahrenheit类型可以使用与float64相同的算术操作符
	fmt.Printf("%g\n", BoilingC-FreezlingC)
	boilingC := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingC-CtoF(FreezlingC))
	//fmt.Printf("%g\n", boilingC-FreezlingC) // 编译失败，类型不匹配
}
