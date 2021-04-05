package chunk

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func TestBext(t *testing.T) {
	chunk := &Bext{}
	headerSize := uint32(12)
	header := EncodeChunkHeader(CreateFourCC(BEXTID), headerSize, binary.LittleEndian)
	chunk.Header = header

	assertEqual(t, chunk.ID(), BEXTID, "id")
	assertEqual(t, chunk.Size(), headerSize, "size")

	description := createTestString(256)
	chunk.SetDescription(description)
	assertEqual(t, chunk.Description(), description, "description")

	description = createTestString(12)
	chunk.SetDescription(description)
	assertEqual(t, chunk.Description(), description, "description")

	originator := createTestString(32)
	chunk.SetOriginator(originator)
	assertEqual(t, chunk.Originator(), originator, "originator")

	originatorReference := createTestString(32)
	chunk.SetOriginatorReference(originatorReference)
	assertEqual(t, chunk.OriginatorReference(), originatorReference, "originatorReference")

	originationDate := createTestString(10)
	chunk.SetOriginationDate(originatorReference)
	assertEqual(t, chunk.OriginationDate(), originationDate, "originationDate")

	originationTime := createTestString(8)
	chunk.SetOriginationTime(originatorReference)
	assertEqual(t, chunk.OriginationTime(), originationTime, "originationTime")

	timeReference := uint64(^uint(0))
	chunk.SetTimeReference(timeReference)
	assertEqual(t, chunk.TimeReference(), timeReference, "timeReference")

	version := uint16(1)
	chunk.SetVersion(version)
	assertEqual(t, chunk.Version(), version, "version")

	umid := [64]uint8{}

	for i := 0; i < len(umid); i++ {
		umid[i] = uint8(i + 1)
	}

	chunk.SetUMID(umid)
	assertEqual(t, chunk.UMID(), umid, "umid")

	var loudnessValue uint16 = 1000
	chunk.SetLoudnessValue(loudnessValue)
	assertEqual(t, chunk.LoudnessValue(), loudnessValue, "loudnessValue")

	var loudnessRange uint16 = 1001
	chunk.SetLoudnessRange(loudnessRange)
	assertEqual(t, chunk.LoudnessRange(), loudnessRange, "loudnessRange")

	var maxTruePeakLevel uint16 = 1002
	chunk.SetMaxTruePeakLevel(maxTruePeakLevel)
	assertEqual(t, chunk.MaxTruePeakLevel(), maxTruePeakLevel, "maxTruePeakLevel")

	var maxMomentaryLoudness uint16 = 1003
	chunk.SetMaxMomentaryLoudness(maxMomentaryLoudness)
	assertEqual(t, chunk.MaxMomentaryLoudness(), maxMomentaryLoudness, "maxMomentaryLoudness")

	var maxShortTermLoudness uint16 = 1004
	chunk.SetMaxShortTermLoudness(maxShortTermLoudness)
	assertEqual(t, chunk.MaxShortTermLoudness(), maxShortTermLoudness, "maxShortTermLoudness")

	codingHistory := createTestString(8)
	chunk.SetCodingHistory(codingHistory)
	assertEqual(t, chunk.CodingHistory(), codingHistory, "codingHistory")

	bytes := chunk.Bytes()
	bext, err := DecodeBextChunk(bytes)
	assertEqual(t, err, nil, "err")
	assertEqual(t, bext.CodingHistory(), codingHistory, "codingHistory")
}

func TestBextBytes(t *testing.T) {
	chunk := &Bext{}
	bytes := chunk.Bytes()

	assertNotNil(t, bytes, "bytes")
	assertEqual(t, len(bytes), 610, "bext chunk length")

	chunk.SetVersion(0)
	history := "coding history"
	chunk.SetCodingHistory(history)
	chunk.SetLoudnessRange(10)
	chunk.SetLoudnessValue(11)
	chunk.SetMaxTruePeakLevel(12)
	chunk.SetMaxMomentaryLoudness(13)
	chunk.SetMaxShortTermLoudness(14)
	umid := [64]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	chunk.SetUMID(umid)
	bytes = chunk.Bytes()
	bext, err := DecodeBextChunk(bytes)

	assertNil(t, err, "err: DecodeBextChunk")
	assertEqual(t, bext.Version(), uint16(0), "Version")
	assertEqual(t, bext.UMID(), [64]byte{}, "UMID")
	assertEqual(t, bext.LoudnessRange(), uint16(0), "LoudnessRange")
	assertEqual(t, bext.LoudnessValue(), uint16(0), "LoudnessValue")
	assertEqual(t, bext.MaxTruePeakLevel(), uint16(0), "MaxTruePeakLevel")
	assertEqual(t, bext.MaxMomentaryLoudness(), uint16(0), "MaxMomentaryLoudness")
	assertEqual(t, bext.MaxShortTermLoudness(), uint16(0), "MaxShortTermLoudness")
	assertEqual(t, bext.CodingHistory(), history, "CodingHistory")

	chunk.SetVersion(1)
	bytes = chunk.Bytes()
	bext, err = DecodeBextChunk(bytes)

	assertNil(t, err, "err: DecodeBextChunk")
	assertEqual(t, bext.Version(), uint16(1), "Version")
	assertEqual(t, bext.UMID(), umid, "UMID")
	assertEqual(t, bext.LoudnessRange(), uint16(0), "LoudnessRange")
	assertEqual(t, bext.LoudnessValue(), uint16(0), "LoudnessValue")
	assertEqual(t, bext.MaxTruePeakLevel(), uint16(0), "MaxTruePeakLevel")
	assertEqual(t, bext.MaxMomentaryLoudness(), uint16(0), "MaxMomentaryLoudness")
	assertEqual(t, bext.MaxShortTermLoudness(), uint16(0), "MaxShortTermLoudness")
	assertEqual(t, bext.CodingHistory(), history, "CodingHistory")

	chunk.SetVersion(2)
	bytes = chunk.Bytes()
	bext, err = DecodeBextChunk(bytes)

	assertNil(t, err, "err: DecodeBextChunk")
	assertEqual(t, bext.Version(), uint16(2), "Version")
	assertEqual(t, bext.UMID(), umid, "UMID")
	assertEqual(t, bext.LoudnessRange(), uint16(10), "LoudnessRange")
	assertEqual(t, bext.LoudnessValue(), uint16(11), "LoudnessValue")
	assertEqual(t, bext.MaxTruePeakLevel(), uint16(12), "MaxTruePeakLevel")
	assertEqual(t, bext.MaxMomentaryLoudness(), uint16(13), "MaxMomentaryLoudness")
	assertEqual(t, bext.MaxShortTermLoudness(), uint16(14), "MaxShortTermLoudness")
	assertEqual(t, bext.CodingHistory(), history, "CodingHistory")
}

func createTestString(size int) string {
	var bytes = make([]byte, size)

	for i := 0; i < len(bytes); i++ {
		bytes[i] = 65
	}

	return string(bytes)
}

func assertEqual(t *testing.T, a interface{}, b interface{}, testObject string) {
	if a == b {
		return
	}

	message := fmt.Sprintf("%s is %v, want %v", testObject, a, b)
	t.Fatal(message)
}

func assertNil(t *testing.T, a interface{}, testObject string) {
	if reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil() {
		return
	}

	if a == nil {
		return
	}

	message := fmt.Sprintf("%s is not nil", testObject)
	t.Fatal(message)
}

func assertNotNil(t *testing.T, a interface{}, testObject string) {
	if a != nil {
		return
	}

	message := fmt.Sprintf("%s is nil", testObject)
	t.Fatal(message)
}
