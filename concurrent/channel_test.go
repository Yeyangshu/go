package concurrent

import "testing"

// 传统的线程之间可以通过共享内存进行数据交互，不同的线程共享内存的同步问题需要使用锁来解决，这样会导致性能低下。
// Go语言中提倡使用channel的方式代替共享内存。换言之，Go语言主张通过数据传递来实现共享内存，而不是通过共享内存来实现数据传递。

// TestCreateChannel 创建channel
func TestCreateChannel(t *testing.T) {
	//ch1 := make(chan int)         // 创建一个整型类型channel
	//ch2 := make(chan interface{}) // 创建一个空接口类型channel

	//type Equip struct{}
	//ch3 := make(chan *Equip) // 创建一个Equip指针类型channel
}
