package chunk

import (
	"testing"
)

func TestTerminate(t *testing.T) {
	maxlength := 4
	val := "test"

	res := terminate(val, maxlength)
	assertEqual(t, len(res), maxlength, "length after terminate")
	assertEqual(t, res, val, "value after terminate")

	res = terminate(val, maxlength-1)
	assertEqual(t, len(res), maxlength-1, "length after terminate")
	assertEqual(t, res, val[:maxlength-1], "value after terminate")

	res = terminate(val[:1], maxlength)
	assertEqual(t, len(res), len(val[:1])+1, "length after terminate")
	assertEqual(t, res, val[:1]+"\x00", "value after terminate")
}
