package concurrent

import (
	"fmt"
	"testing"
	"time"
)

// TestNormalFuncCreateGoroutine 普通函数创建Goroutine
// 普通函数创建goroutine，在函数或方法前加上关键字go，将会运行一个新的Goroutine。
// 使用Go关键字创建Goroutine时，被调用的函数往往我没有返回值，如果有返回值也会被忽略。
// 如果需要在Goroutine中返回数据，必须使用channel。
func TestNormalFuncCreateGoroutine(t *testing.T) {
	// 子Goroutine
	go hello()
	// 如果主程序执行完成，主Goroutine终止，程序终止，其他的Goroutine不再运行。
	// 如果主Goroutine比子Goroutine先终止，运行结果不会打印"Hello world goroutine"
	fmt.Println("main function")
}

// TestNormalFuncCreateGoroutineSleep 普通函数创建Goroutine，sleep子Goroutine
func TestNormalFuncCreateGoroutineSleep(t *testing.T) {
	// 子Goroutine
	go hello()

	time.Sleep(1 * time.Millisecond)

	// 主程序Sleep，子Goroutine执行完成
	fmt.Println("main function")
}

func hello() {
	fmt.Println("Hello world goroutine")
}

// TestAnonymousFunctionCreateGoroutine 匿名函数创建Goroutine
func TestAnonymousFunctionCreateGoroutine(t *testing.T) {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}

// TestStartMultiGoroutine 启动多个Goroutine
func TestStartMultiGoroutine(t *testing.T) {
	// 多个Goroutine随机调度
	go printNum()
	go printLetter()

	time.Sleep(3 * time.Second)
	fmt.Println("\n main over")
}

func printNum() {
	for i := 0; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetter() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
