package binp

import (
	"math"
	"unsafe"
)

func ntohq(x uint64) uint64 {
	return ((uint64(ntohl(uint32((x) >> 32)))) | (uint64(ntohl(uint32(x))) << 32))
}

// Take the first 8 bigendian bytes of a buffer and return them.
func B64(b []byte) uint64 {
	if len(b) < 8 {
		panic("B64: too short buffer")
	}
	return ntohq(*(*uint64)(unsafe.Pointer(&b[0])))
}

// Poke 8 bigendian bytes into the buffer.
func PutB64(b []byte, v uint64) {
	if len(b) < 8 {
		panic("PutB64: too short buffer")
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = ntohq(v)
}

// Take the first 4 bigendian bytes of a buffer and return them.
func B32(b []byte) uint32 {
	if len(b) < 4 {
		panic("B32: too short buffer")
	}
	return ntohl(*(*uint32)(unsafe.Pointer(&b[0])))
}

// Poke 4 bigendian bytes into the buffer.
func PutB32(b []byte, v uint32) {
	if len(b) < 4 {
		panic("PutB32: too short buffer")
	}
	*(*uint32)(unsafe.Pointer(&b[0])) = ntohl(v)
}

// Take the first 2 bigendian bytes of a buffer and return them.
func B16(b []byte) uint16 {
	if len(b) < 2 {
		panic("B16: too short buffer")
	}
	return ntohs(*(*uint16)(unsafe.Pointer(&b[0])))
}

// Poke 2 bigendian bytes into the buffer.
func PutB16(b []byte, v uint16) {
	if len(b) < 2 {
		panic("PutB16: too short buffer")
	}
	*(*uint16)(unsafe.Pointer(&b[0])) = ntohs(v)
}

// PutBF32 pokes a float32.
func PutBF32(b []byte, v float32) {
	PutB32(b, math.Float32bits(v))
}

// BF32 peeks a float32
func BF32(b []byte) float32 {
	return math.Float32frombits(B32(b))
}

// PutBF64 pokes a float64.
func PutBF64(b []byte, v float64) {
	PutB64(b, math.Float64bits(v))
}

// BF64 peeks a float64
func BF64(b []byte) float64 {
	return math.Float64frombits(B64(b))
}
