# 协程的概念
协程是编译器级的，线程是操作系统级别的。协程不被操作系统内核管理，而完全由程序控制，因此没有线程切换的开销。

Go 语言中的协程叫做 Goroutine