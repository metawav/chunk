package chunk

import "testing"

func TestCreateFourCC(t *testing.T) {
	fourCC := CreateFourCC("")
	assertEqual(t, len(fourCC), 4, "fouCC length")

	fourCC = CreateFourCC("TEST")
	assertEqual(t, fourCC.String(), "TEST", "String")

	fourCC = CreateFourCC("TEST1")
	assertEqual(t, fourCC.String(), "TEST", "String")
}
