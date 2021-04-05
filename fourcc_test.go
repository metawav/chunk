package chunk

import "testing"

func TestCreateFourCC(t *testing.T) {
	fourCC := CreateFourCC("")

	if len(fourCC) != 4 {
		t.Errorf("foucc length is %d, want %d", len(fourCC), 4)
	}

	fourCC = CreateFourCC("TEST")

	if fourCC.String() != "TEST" {
		t.Errorf("foucc is %s, want %s", fourCC, "TEST")
	}

	fourCC = CreateFourCC("TEST1")

	if fourCC.String() != "TEST" {
		t.Errorf("foucc is %s, want %s", fourCC, "TEST")
	}
}
