package utils

type NilInt struct {
	Value int
	Null  bool
}

func (n *NilInt) GetValue() interface{} {
	if n.Null {
		return nil
	}
	return n.Value
}

func NewInt(x int) NilInt {
	return NilInt{x, false}
}

func NewNil() NilInt {
	return NilInt{0, true}
}
