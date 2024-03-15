package codec

import "io"

type Header struct {
	ServiceMethod string // format: "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

// Codec for encoding and decoding
type Codec interface {
	io.Closer // close the connection
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}
type NewCoderFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCoderFuncMap map[Type]NewCoderFunc

func init() {
	NewCoderFuncMap = make(map[Type]NewCoderFunc)
	NewCoderFuncMap[GobType] = NewGobCodec

}
