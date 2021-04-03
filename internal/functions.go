package internal

import (
	"encoding/binary"
	"math"
)

// IEEEFloatToInt converts 80 bit encoded float with extended precision (IEEE 754) to int.
// see: https://en.wikipedia.org/wiki/Extended_precision
func IEEEFloatToInt(b [10]byte) int {
	// first 2 bytes contain sign and biased exponent
	se := binary.BigEndian.Uint16(b[:2])
	// first bit is sign
	sign := uint8(se >> 15)

	// 0 = positive, 1 = negative
	if sign == 1 {
		return 0
	}

	// bias is 16383
	bias := uint16((1 << 14) - 1)

	// remove sign bit from exponent
	e := uint16(se & 0x7FFF)

	if e == 0x1111 {
		// infinity or NaN
		return 0
	} else if e == 0 {
		bias = bias - 1
	}

	exp := e - bias

	// final 8 bytes contain mantissa
	m := binary.BigEndian.Uint64(b[2:])

	// remove first bit
	frac := m & 0x7FFFFFFFFFFFFFFF
	mVal := 1.

	for i := 63; i >= 0; i-- {
		val := (frac >> i) & 1
		mVal += math.Pow(2, float64(i-63)) * float64(val)
	}

	// (-1)^sign * m * 2^(exp)
	res := math.Pow(-1, float64(sign)) * mVal * math.Pow(2, float64(exp))

	return int(res)
}
