/*
The bigendian package provides parsing and printing of binary bigendian values. 
It is particularly suitable for parsing network protocols and creating packets.

For instance printing:

	p := bigendian.NewPrinter().Byte(1).Bytes(rand(20))
	p.U32String("foo").U32String("bar")
	p.Byte(0).U32(0)
	return p.Out()


And parsing

	var ptype byte
	var cookie []byte
	var kex, shk string
	var follows byte
	var reserved uint32

	p := bigendian.NewParser(b).Byte(&ptype).NBytes(16, &cookie)
	p.U32String(&kex).U32String(&shk)
	p.Byte(&follows).U32(&reserved).End()

*/
package binp

import "unsafe"

func ntohq(x uint64) uint64 {
	return ((uint64(ntohl(uint32((x)>>32))))|((uint64(ntohl(uint32(x)))<<32)))
}

// Take the first 8 bigendian bytes of a buffer and return them.
func U64(b []byte) uint64 {
	if len(b) < 8 {
		panic("U64: too short buffer")
	}
	return ntohq(*(*uint64)(unsafe.Pointer(&b[0])))
}

// Poke 8 bigendian bytes into the buffer.
func PutU64(b []byte, v uint64) {
	if len(b) < 8 {
		panic("PutU64: too short buffer")
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = ntohq(v)
}


// Take the first 4 bigendian bytes of a buffer and return them.
func U32(b []byte) uint32 {
	if len(b) < 4 {
		panic("U32: too short buffer")
	}
	return ntohl(*(*uint32)(unsafe.Pointer(&b[0])))
}

// Poke 4 bigendian bytes into the buffer.
func PutU32(b []byte, v uint32) {
	if len(b) < 4 {
		panic("PutU32: too short buffer")
	}
	*(*uint32)(unsafe.Pointer(&b[0])) = ntohl(v)
}


// Take the first 2 bigendian bytes of a buffer and return them.
func U16(b []byte) uint16 {
	if len(b) < 2 {
		panic("U16: too short buffer")
	}
	return ntohs(*(*uint16)(unsafe.Pointer(&b[0])))
}

// Poke 2 bigendian bytes into the buffer.
func PutU16(b []byte, v uint16) {
	if len(b) < 2 {
		panic("PutU16: too short buffer")
	}
	*(*uint16)(unsafe.Pointer(&b[0])) = ntohs(v)
}
