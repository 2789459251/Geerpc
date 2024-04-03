package Geerpc

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo int
type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num2 + args.Num1
	return nil
}
func _assert(condution bool, msg string, v ...interface{}) {
	if !condution {
		panic(fmt.Sprintf("assertion faild:"+msg, v...))
	}
}
func TestNewService(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	_assert(len(s.method) == 1, "wrong service Method,expect 1,but got %d", len(s.method))
	mType := s.method["Sum"]
	_assert(mType != nil, "wrong Method", "sum should not nil")
}

func TestMethodType_NumCalls(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	mType := s.method["Sum"]

	argv := mType.newArgv()
	replyv := mType.newReplyv()

	argv.Set(reflect.ValueOf(Args{
		Num1: 1,
		Num2: 3,
	}))

	err := s.call(mType, argv, replyv)
	_assert(err == nil && *replyv.Interface().(*int) == 4 && mType.NumCalls() == 1, "failed to call Foo.sum")
}
