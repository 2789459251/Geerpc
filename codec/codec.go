package codec

import "io"

type Header struct {
	ServiceMethod string //服务名与方法名：服务.方法
	Seq           uint64 //请求的序列号，区分不同请求
	Error         string
}
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
type NewCodecFunc func(io.ReadWriteCloser) Codec
type Type string

/*
gob是Golang包自带的一个数据结构序列化的编码/解码工具。 编码使用Encoder，解码使用Decoder。
一种典型的应用场景就是RPC(remote procedure calls)。
gob和json的pack之类的方法一样，由发送端使用Encoder对数据结构进行编码。
*/
const (
	GobType  Type = "application/job"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc //值是方法，不是方法的返回值

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
