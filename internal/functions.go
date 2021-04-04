package internal

import (
	"encoding/binary"
	"math"
)

// IEEEFloatToInt converts 80 bit encoded float with extended precision (IEEE 754) to int.
// For negative values, infinity and NaN (+)0 is returned.
// See: https://en.wikipedia.org/wiki/Extended_precision
func IEEEFloatToInt(b [10]byte) uint {
	// first 2 bytes contain sign and biased exponent
	se := binary.BigEndian.Uint16(b[:2])

	// if all bits of se are zero just return 0
	if se == 0 {
		return 0
	}

	// first bit is sign
	sign := uint8(se >> 15)

	// 0 = positive, 1 = negative
	if sign == 1 {
		// return 0 for negative values
		return 0
	}

	// bias is 16383
	bias := uint16((1 << 14) - 1)

	// remove sign bit from biased exponent
	e := uint16(se & 0x7FFF)

	if e == 0x1111 {
		// infinity or NaN
		return 0
	} else if e == 0 {
		bias = bias - 1
	}

	// remove bias
	exp := e - bias

	// final 8 bytes contain mantissa
	m := binary.BigEndian.Uint64(b[2:])

	// remove first bit from mantissa, as in 80-bit representation this is the integer part
	frac := m & 0x7FFFFFFFFFFFFFFF
	mVal := 1.

	for i := 63; i >= 0; i-- {
		val := (frac >> i) & 1
		mVal += math.Pow(2, float64(i-63)) * float64(val)
	}

	// (-1)^sign * m * 2^(exp)
	res := math.Pow(-1, float64(sign)) * mVal * math.Pow(2, float64(exp))

	return uint(res)
}

// IntToIEEE is a simplified conversion of positive int values to 80 bit encoded float with extended precision (IEE 754).
func IntToIEEE(v uint) [10]byte {
	var res [10]byte

	if v == 0 {
		return res
	}

	// 8 bytes mantissa
	q := v
	var i uint16
	var m, temp uint64

	for q > 0 {
		x := uint64(q % 2)
		temp = temp | x<<i
		q = q / 2
		i++
	}

	m = m | temp<<(63-i+1)
	binary.BigEndian.PutUint64(res[2:], m)

	// bias is 16383
	bias := uint16((1 << 14) - 1)
	// biased exponent
	exp := i - 1 + bias

	// first bit is sign
	var sign uint16

	// first 2 bytes contain sign and biased exponent
	var se uint16 = sign<<15 | exp
	binary.BigEndian.PutUint16(res[0:2], se)

	return res
}
