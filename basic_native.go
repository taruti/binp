/*
The binp package provides parsing and printing of binary values.
It is particularly suitable for parsing irregular data structures provided by e.g. kernel.

*/
package binp

import "unsafe"

// Take the first 8 native endian bytes of a buffer and return them.
func N64(b []byte) uint64 {
	if len(b) < 8 {
		panic("N64: too short buffer")
	}
	return *(*uint64)(unsafe.Pointer(&b[0]))
}

// Poke 8 native endian bytes into the buffer.
func PutN64(b []byte, v uint64) {
	if len(b) < 8 {
		panic("PutN64: too short buffer")
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = v
}

// Take the first 4 native endian bytes of a buffer and return them.
func N32(b []byte) uint32 {
	if len(b) < 4 {
		panic("N32: too short buffer")
	}
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

// Poke 4 native endian bytes into the buffer.
func PutN32(b []byte, v uint32) {
	if len(b) < 4 {
		panic("PutN32: too short buffer")
	}
	*(*uint32)(unsafe.Pointer(&b[0])) = v
}

// Take the first 2 native endian bytes of a buffer and return them.
func N16(b []byte) uint16 {
	if len(b) < 2 {
		panic("N16: too short buffer")
	}
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

// Poke 2 native endian bytes into the buffer.
func PutN16(b []byte, v uint16) {
	if len(b) < 2 {
		panic("PutN16: too short buffer")
	}
	*(*uint16)(unsafe.Pointer(&b[0])) = v
}
