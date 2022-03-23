package rtda

// Thread 线程
type Thread struct {
	pc       int       // pc计数器
	jvmStack *JVMStack // java虚拟机栈
}
