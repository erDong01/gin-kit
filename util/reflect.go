package util

import (
	"fmt"
	"strconv"
)

// I interface{}类型转换
type I interface {
	Bytes() []byte
	String() string
	Int64() int64
	Int() int
	Uint() uint
	Float64() float64
	Uint64() uint64
	Bool() bool
}

func Reflect(v interface{}) I {
	return &t{i: v}
}

type t struct {
	i interface{}
}

func (t *t) String() string {
	if t.i == nil {
		return ""
	}
	return fmt.Sprint(t.i)
}

func (t *t) Bytes() []byte {
	return []byte(t.String())
}

func (t *t) Int64() int64 {
	i, err := strconv.ParseInt(t.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (t *t) Int() int {
	return int(t.Int64())
}

func (t *t) Uint() uint {
	return uint(t.Uint64())
}

func (t *t) Uint64() uint64 {
	i, err := strconv.ParseUint(t.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (t *t) Float64() float64 {
	f, err := strconv.ParseFloat(t.String(), 64)
	if err != nil {
		return 0
	}
	return f
}

func (t *t) Bool() bool {
	b, err := strconv.ParseBool(t.String())
	if err != nil {
		return false
	}
	return b
}
