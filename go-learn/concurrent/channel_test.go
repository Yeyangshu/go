package concurrent

import (
	"fmt"
	"testing"
)

// 传统的线程之间可以通过共享内存进行数据交互，不同的线程共享内存的同步问题需要使用锁来解决，这样会导致性能低下。
// Go语言中提倡使用channel的方式代替共享内存。换言之，Go语言主张通过数据传递来实现共享内存，而不是通过共享内存来实现数据传递。

// TestCreateChannel 创建channel
func TestCreateChannel(t *testing.T) {
	//ch1 := make(chan int)         // 创建一个整型类型channel
	//ch2 := make(chan interface{}) // 创建一个空接口类型channel

	//type Equip programstructure{}
	//ch3 := make(chan *Equip) // 创建一个Equip指针类型channel
}

// TestReceiveChannelData 接收channel数据
func TestReceiveChannelData(t *testing.T) {
	ch1 := make(chan string)
	go sendData(ch1)

	for {
		data := <-ch1
		// 如果通道关闭，通道中传输的数据为各类型的零值
		if data == "" {
			// for循环必须使用break判断来终止循环
			break
		}
		fmt.Println("从通道中读取数据：", data)
	}

	for {
		data, ok := <-ch1
		// 如果通道关闭，通道中传输的数据为各类型的零值
		if !ok {
			break
		}
		fmt.Println("从通道中读取数据：", data)
	}

	for value := range ch1 {
		fmt.Println("从通道中读取数据：", value)
	}
}

func sendData(ch1 chan string) {
	defer close(ch1)

	for i := 0; i < 3; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}

	fmt.Sprintln("发送数据完毕")
}

// TestChannelBlock channel阻塞
func TestChannelBlock(t *testing.T) {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)
	ch2 := make(chan bool)
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("子goroutine取到数值：", data)
			ch2 <- true
		}
	}()
	ch1 <- 10
	// 阻塞
	<-ch2
	fmt.Println("main over...")
}

// TestCloseChannel 关闭channel
func TestCloseChannel(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		ch1 <- 100
		ch1 <- 200
		close(ch1)
		// 关闭的channel，无法写入数据
		ch1 <- 10
	}()

	data, ok := <-ch1
	fmt.Println("main取到数值：", data, ok)
	data, ok = <-ch1
	fmt.Println("main取到数值：", data, ok)
	data, ok = <-ch1
	fmt.Println("main取到数值：", data, ok)
	data, ok = <-ch1
	fmt.Println("main取到数值：", data, ok)
}
