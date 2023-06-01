package concurrent

import (
	"fmt"
	"testing"
)

// TestBufferChannel 缓冲通道
func TestBufferChannel(t *testing.T) {
	ch1 := make(chan string, 6)
	go sendData(ch1)
	for data := range ch1 {
		fmt.Println("\t读取数据", data)
	}
}
