package Geerpc

import (
	"reflect"
	"sync/atomic"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

func (m *methodType) NumCalls() uint64 {
	//atomic.LoadUint64 函数会返回 numCalls 当前的值，而不会修改它。
	//这保证了在并发情况下，即使其他 goroutine 在同时修改 numCalls 的值
	//我们也能安全地读取到其当前值，而不会发生数据竞态。
	return atomic.LoadUint64(&m.numCalls)
}
