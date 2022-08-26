package rpc

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"reflect"
	"strings"

	"github.com/erDong01/micro-kit/base"
	"google.golang.org/protobuf/proto"
)

// rpc UnmarshalHead
func UnmarshalHead(buff []byte) (*RpcPacket, RpcHead) {
	nLen := base.Clamp(len(buff), 0, 256)
	return Unmarshal(buff[:nLen])
}
func Unmarshal(buff []byte) (*RpcPacket, RpcHead) {
	rpcPacket := &RpcPacket{}
	proto.Unmarshal(buff, rpcPacket)
	if rpcPacket.RpcHead == nil {
		rpcPacket.RpcHead = &RpcHead{}
	}
	rpcPacket.FuncName = strings.ToLower(rpcPacket.FuncName)
	return rpcPacket, *(*RpcHead)(rpcPacket.RpcHead)
}

// rpc Unmarshal
// pFuncType for RegisterCall func
func UnmarshalBody(rpcPacket *RpcPacket, pFuncType reflect.Type) []interface{} {
	nCurLen := pFuncType.NumIn()
	params := make([]interface{}, nCurLen)
	var dec *gob.Decoder
	if rpcPacket.ArgLen > 0 {
		buf := bytes.NewBuffer(rpcPacket.RpcBody)
		dec = gob.NewDecoder(buf)
	}
	for i := 0; i < nCurLen; i++ {
		if i == 0 {
			params[0] = context.WithValue(context.Background(), "rpcHead", *(*RpcHead)(rpcPacket.RpcHead))
			continue
		}
		if i < int(rpcPacket.ArgLen+1) {
			val := reflect.New(pFuncType.In(i))
			dec.DecodeValue(val)
			params[i] = val.Elem().Interface()
		}
		if rpcPacket.ArgLen == 0 {
			val := reflect.New(pFuncType.In(i).Elem())
			m := val.Interface().(proto.Message)
			proto.Unmarshal(rpcPacket.RpcBody, m)
			params[i] = m
		}
	}
	return params
}

func UnmarshalBodyCall(rpcPacket *RpcPacket, pFuncType reflect.Type) (error, []interface{}) {
	strErr := ""
	nCurLen := pFuncType.NumIn()
	params := make([]interface{}, nCurLen)
	buf := bytes.NewBuffer(rpcPacket.RpcBody)
	dec := gob.NewDecoder(buf)
	dec.Decode(&strErr)
	if strErr != "" {
		return errors.New(strErr), params
	}
	for i := 0; i < nCurLen; i++ {
		if i == 0 {
			params[0] = context.WithValue(context.Background(), "rpcHead", *(*RpcHead)(rpcPacket.RpcHead))
			continue
		}

		val := reflect.New(pFuncType.In(i))
		if i < int(rpcPacket.ArgLen+1) {
			dec.DecodeValue(val)
		}
		params[i] = val.Elem().Interface()
	}
	return nil, params
}
