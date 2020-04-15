package boolcast

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

func CastUnsafe(b bool) byte {
	return *(*byte)(unsafe.Pointer(&b)) & 1
}

func CastConditional(b bool) byte {
	if b {
		return 1
	}
	return 0
}

var _bMap = map[bool]byte{
	true:  1,
	false: 0,
}

func CastMapped(b bool) byte {
	return _bMap[b]
}

var _bCastBuf = bytes.NewBuffer(make([]byte, 1))

func CastStandard(b bool) byte {
	_bCastBuf.Reset()
	_ = binary.Write(_bCastBuf, binary.LittleEndian, b)
	return _bCastBuf.Bytes()[0]
}

func CastFoundationDB(b bool) byte {
	return tuple.Tuple{b}.Pack()[0]
}

func init() {
	fdb.APIVersion(610)
}
