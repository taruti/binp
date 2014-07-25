/*
The native package provides parsing and printing of binary native endian values.
It is particularly suitable for parsing irregular data structures provided by e.g. kernel.

For instance printing:

	p := native.NewPrinter().Byte(1).Bytes(rand(20))
	p.U32String("foo").U32String("bar")
	p.Byte(0).U32(0)
	return p.Out()


And parsing

	var ptype byte
	var cookie []byte
	var kex, shk string
	var follows byte
	var reserved uint32

	p := native.NewParser(b).Byte(&ptype).NBytes(16, &cookie)
	p.U32String(&kex).U32String(&shk)
	p.Byte(&follows).U32(&reserved).End()

*/
package native

import "unsafe"

// Take the first 8 bigendian bytes of a buffer and return them.
func U64(b []byte) uint64 {
	if len(b) < 8 {
		panic("U64: too short buffer")
	}
	return *(*uint64)(unsafe.Pointer(&b[0]))
}

// Poke 8 bigendian bytes into the buffer.
func PutU64(b []byte, v uint64) {
	if len(b) < 8 {
		panic("PutU64: too short buffer")
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = v
}

// Take the first 4 bigendian bytes of a buffer and return them.
func U32(b []byte) uint32 {
	if len(b) < 4 {
		panic("U32: too short buffer")
	}
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

// Poke 4 bigendian bytes into the buffer.
func PutU32(b []byte, v uint32) {
	if len(b) < 4 {
		panic("PutU32: too short buffer")
	}
	*(*uint32)(unsafe.Pointer(&b[0])) = v
}

// Take the first 2 bigendian bytes of a buffer and return them.
func U16(b []byte) uint16 {
	if len(b) < 2 {
		panic("U16: too short buffer")
	}
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

// Poke 2 bigendian bytes into the buffer.
func PutU16(b []byte, v uint16) {
	if len(b) < 2 {
		panic("PutU16: too short buffer")
	}
	*(*uint16)(unsafe.Pointer(&b[0])) = v
}
