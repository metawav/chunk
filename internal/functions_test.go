package internal

import (
	"testing"
)

func TestIntToIEEE(t *testing.T) {
	vals := []uint{0, 1, 8000, 41400, 4800, 96000, 192000, 384000}

	for _, v := range vals {
		r := IntToIEEE(v)
		res := IEEEFloatToInt(r)

		if res != v {
			t.Errorf("result is %d, want%d", res, v)
		}
	}
}
