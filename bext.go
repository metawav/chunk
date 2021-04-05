package chunk

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// Bext is Broadcast Wave Format (BWF) bext chunk 'bext' describing extended information.
// This data structure is compatible with BWF versions 0, 1 and 2.
// Version 0 has 254 bytes reserve.
// Version 1 contains 64 byte of UMID and has 190 bytes reserve.
// Version 2 is like version 1 additionally contains 10 byte loudness information and has 180 bytes reserve.
type Bext struct {
	*Header
	description          [256]byte
	originator           [32]byte
	originatorReference  [32]byte
	originationDate      [10]byte
	originationTime      [8]byte
	timeReferenceLow     uint32
	timeReferenceHigh    uint32
	version              uint16
	uMID0                uint8
	uMID1                uint8
	uMID2                uint8
	uMID3                uint8
	uMID4                uint8
	uMID5                uint8
	uMID6                uint8
	uMID7                uint8
	uMID8                uint8
	uMID9                uint8
	uMID10               uint8
	uMID11               uint8
	uMID12               uint8
	uMID13               uint8
	uMID14               uint8
	uMID15               uint8
	uMID16               uint8
	uMID17               uint8
	uMID18               uint8
	uMID19               uint8
	uMID20               uint8
	uMID21               uint8
	uMID22               uint8
	uMID23               uint8
	uMID24               uint8
	uMID25               uint8
	uMID26               uint8
	uMID27               uint8
	uMID28               uint8
	uMID29               uint8
	uMID30               uint8
	uMID31               uint8
	uMID32               uint8
	uMID33               uint8
	uMID34               uint8
	uMID35               uint8
	uMID36               uint8
	uMID37               uint8
	uMID38               uint8
	uMID39               uint8
	uMID40               uint8
	uMID41               uint8
	uMID42               uint8
	uMID43               uint8
	uMID44               uint8
	uMID45               uint8
	uMID46               uint8
	uMID47               uint8
	uMID48               uint8
	uMID49               uint8
	uMID50               uint8
	uMID51               uint8
	uMID52               uint8
	uMID53               uint8
	uMID54               uint8
	uMID55               uint8
	uMID56               uint8
	uMID57               uint8
	uMID58               uint8
	uMID59               uint8
	uMID60               uint8
	uMID61               uint8
	uMID62               uint8
	uMID63               uint8
	loudnessValue        uint16
	loudnessRange        uint16
	maxTruePeakLevel     uint16
	maxMomentaryLoudness uint16
	maxShortTermLoudness uint16
	reserved             [180]byte
	codingHistory        []byte
}

// Description is a free description.
// Max. 256 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) Description() string {
	return nullTermToString(b.description[:])
}

// SetDescription sets a free description.
// Max. 256 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) SetDescription(value string) {
	b.description = [256]byte{}
	copy(b.description[:], terminate(value, len(b.description)))
}

// Originator contains the name of the originator.
// Max. 32 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) Originator() string {
	return nullTermToString(b.originator[:])
}

// SetOriginator sets the name of the originator.
// Max. 32 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) SetOriginator(value string) {
	b.originator = [32]byte{}
	copy(b.originator[:], terminate(value, len(b.originator)))
}

// OriginatorReference
// Max. 32 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) OriginatorReference() string {
	return nullTermToString(b.originatorReference[:])
}

// SetOriginatorReference
// Max. 32 characters and null terminated if shorter.
// (since version 0)
func (b *Bext) SetOriginatorReference(value string) {
	b.originatorReference = [32]byte{}
	copy(b.originatorReference[:], terminate(value, len(b.originatorReference)))
}

// OriginationDate
// 10 characters YYYY-MM-DD
// (since version 0)
func (b *Bext) OriginationDate() string {
	return string(b.originationDate[:])
}

// SetOriginatorDate
// 10 characters YYYY-MM-DD
// (since version 0)
func (b *Bext) SetOriginationDate(value string) {
	b.originationDate = [10]byte{}
	copy(b.originationDate[:], value)
}

// OriginationTime
// 8 characters HH-MM-SS
// Valid separators are '-', '_', ':', ' ', '.'
// (since version 0)
func (b *Bext) OriginationTime() string {
	return string(b.originationTime[:])
}

// SetOriginatorTime
// 8 characters HH-MM-SS
// Valid separators are '-', '_', ':', ' ', '.'
// (since version 0)
func (b *Bext) SetOriginationTime(value string) {
	b.originationTime = [8]byte{}
	copy(b.originationTime[:], value)
}

// TimeReference is the timecode.
// (since version 0)
func (b *Bext) TimeReference() uint64 {
	var value uint64
	value += uint64(b.timeReferenceLow) << 32
	value += uint64(b.timeReferenceHigh)

	return value
}

// SetTimeReference sets the timecode.
// (since version 0)
func (b *Bext) SetTimeReference(value uint64) {
	b.timeReferenceLow = uint32(value >> 32)
	b.timeReferenceHigh = uint32(value)
}

// Version is the version of Broadcast Wave Format (BWF).
// (since version 0)
func (b *Bext) Version() uint16 {
	return b.version
}

// SetVersion sets the version of Broadcast Wave Format (BWF).
// (since version 0)
func (b *Bext) SetVersion(value uint16) {
	b.version = value
}

// UMID is a Unique Material Identifier (UMID) and is standardized in SMPTE 330M.
// (since version 1)
func (b *Bext) UMID() [64]uint8 {
	umid := [64]uint8{b.uMID0, b.uMID1, b.uMID2, b.uMID3, b.uMID4, b.uMID5, b.uMID6, b.uMID7,
		b.uMID8, b.uMID9, b.uMID10, b.uMID11, b.uMID12, b.uMID13, b.uMID14, b.uMID15,
		b.uMID16, b.uMID17, b.uMID18, b.uMID19, b.uMID20, b.uMID21, b.uMID22, b.uMID23,
		b.uMID24, b.uMID25, b.uMID26, b.uMID27, b.uMID28, b.uMID29, b.uMID30, b.uMID31,
		b.uMID32, b.uMID33, b.uMID34, b.uMID35, b.uMID36, b.uMID37, b.uMID38, b.uMID39,
		b.uMID40, b.uMID41, b.uMID42, b.uMID43, b.uMID44, b.uMID45, b.uMID46, b.uMID47,
		b.uMID48, b.uMID49, b.uMID50, b.uMID51, b.uMID52, b.uMID53, b.uMID54, b.uMID55,
		b.uMID56, b.uMID57, b.uMID58, b.uMID59, b.uMID60, b.uMID61, b.uMID62, b.uMID63}

	return umid
}

// SetUMID sets a Unique Material Identifier (UMID) and is standardized in SMPTE 330M.
// (since version 1)
func (b *Bext) SetUMID(value [64]uint8) {
	fields := [64]*uint8{&b.uMID0, &b.uMID1, &b.uMID2, &b.uMID3, &b.uMID4, &b.uMID5, &b.uMID6, &b.uMID7,
		&b.uMID8, &b.uMID9, &b.uMID10, &b.uMID11, &b.uMID12, &b.uMID13, &b.uMID14, &b.uMID15,
		&b.uMID16, &b.uMID17, &b.uMID18, &b.uMID19, &b.uMID20, &b.uMID21, &b.uMID22, &b.uMID23,
		&b.uMID24, &b.uMID25, &b.uMID26, &b.uMID27, &b.uMID28, &b.uMID29, &b.uMID30, &b.uMID31,
		&b.uMID32, &b.uMID33, &b.uMID34, &b.uMID35, &b.uMID36, &b.uMID37, &b.uMID38, &b.uMID39,
		&b.uMID40, &b.uMID41, &b.uMID42, &b.uMID43, &b.uMID44, &b.uMID45, &b.uMID46, &b.uMID47,
		&b.uMID48, &b.uMID49, &b.uMID50, &b.uMID51, &b.uMID52, &b.uMID53, &b.uMID54, &b.uMID55,
		&b.uMID56, &b.uMID57, &b.uMID58, &b.uMID59, &b.uMID60, &b.uMID61, &b.uMID62, &b.uMID63}

	for i, v := range value {
		*fields[i] = v
	}
}

// LoudnessValue in LUFS.
// (since version 2)
func (b *Bext) LoudnessValue() uint16 {
	return b.loudnessValue
}

// SetLoudnessValue
// (since version 2)
func (b *Bext) SetLoudnessValue(value uint16) {
	b.loudnessValue = value
}

// LoudnessRange in LU.
// (since version 2)
func (b *Bext) LoudnessRange() uint16 {
	return b.loudnessRange
}

// SetLoudnessRange
// (since version 2)
func (b *Bext) SetLoudnessRange(value uint16) {
	b.loudnessRange = value
}

// MaxTruePeakLevel in dBTP.
// (since version 2)
func (b *Bext) MaxTruePeakLevel() uint16 {
	return b.maxTruePeakLevel
}

// SetMaxTruePeakLevel
// (since version 2)
func (b *Bext) SetMaxTruePeakLevel(value uint16) {
	b.maxTruePeakLevel = value
}

// MaxMomentaryLoudness in LUFS.
// (since version 2)
func (b *Bext) MaxMomentaryLoudness() uint16 {
	return b.maxMomentaryLoudness
}

// SetMaxMomentaryLoudness
// (since version 2)
func (b *Bext) SetMaxMomentaryLoudness(value uint16) {
	b.maxMomentaryLoudness = value
}

// MaxShortTermLoudness in LUFS.
// (since version 2)
func (b *Bext) MaxShortTermLoudness() uint16 {
	return b.maxShortTermLoudness
}

// SetMaxShortTermLoudness
// (since version 2)
func (b *Bext) SetMaxShortTermLoudness(value uint16) {
	b.maxShortTermLoudness = value
}

// CodingHistory CR/LF terminated strings containing all applied processing information.
// (since version 0)
func (b *Bext) CodingHistory() string {
	return nullTermToString(b.codingHistory[:])
}

// SetCodingHistory
// (since version 0)
func (b *Bext) SetCodingHistory(value string) {
	b.codingHistory = []byte{}
	b.codingHistory = []byte(value)
}

// String returns string represensation of chunk.
func (b *Bext) String() string {
	return fmt.Sprintf("Description: %s\nOriginator: %s\nOriginator Reference: %s\nOrigination Date: %s\nOrigination Time: %s\nTime Reference: %d\nVersion: %d\nUMID: %08b\nLoudness Value: %d\nLoudness Range: %d\nMax True Peak Level: %d\nMax Momentary Loudness: %d\nMax Short Term Loudness: %d\nCoding History: %s",
		b.Description(), b.Originator(), b.OriginatorReference(), b.OriginationDate(), b.OriginationTime(), b.TimeReference(), b.Version(), b.UMID(), b.LoudnessValue(), b.LoudnessRange(), b.MaxTruePeakLevel(), b.MaxMomentaryLoudness(), b.MaxShortTermLoudness(), b.CodingHistory())
}

// Bytes converts BEXT chunk to byte array. A new Header with id 'bext' is created.
// When Bext is converted to byte array, a check for version is applied and fields
// not compatible with version are ignored.
//
// Version 0: All fields but UMID and loudness fields (254 byte reserve)
// Version 1: All fields but loudness fields (190 byte reserve)
// Version 2: All fields (180 byte reserve)
//
// Header size is set to real data size. A minimum amount of 610 bytes is returned.
// chunk header - 8 bytes
// data - 602 bytes
// coding history - not restricted amount of bytes
//
// A padding byte is added if size is odd. This optional byte is not reflected in size.
func (b *Bext) Bytes() []byte {
	byteOrder := binary.LittleEndian
	data := make([]byte, 602)
	copy(data[:256], b.description[:])
	copy(data[256:288], b.originator[:])
	copy(data[288:320], b.originatorReference[:])
	copy(data[320:330], b.originationDate[:])
	copy(data[330:338], b.originationTime[:])
	byteOrder.PutUint32(data[338:342], b.timeReferenceLow)
	byteOrder.PutUint32(data[342:346], b.timeReferenceHigh)
	byteOrder.PutUint16(data[346:348], b.version)

	umidFields := []uint8{b.uMID0, b.uMID1, b.uMID2, b.uMID3, b.uMID4, b.uMID5, b.uMID6, b.uMID7, b.uMID8, b.uMID9,
		b.uMID10, b.uMID11, b.uMID12, b.uMID13, b.uMID14, b.uMID15, b.uMID16, b.uMID17, b.uMID18, b.uMID19,
		b.uMID20, b.uMID21, b.uMID22, b.uMID23, b.uMID24, b.uMID25, b.uMID26, b.uMID27, b.uMID28, b.uMID29,
		b.uMID30, b.uMID31, b.uMID32, b.uMID33, b.uMID34, b.uMID35, b.uMID36, b.uMID37, b.uMID38, b.uMID39,
		b.uMID40, b.uMID41, b.uMID42, b.uMID43, b.uMID44, b.uMID45, b.uMID46, b.uMID47, b.uMID48, b.uMID49,
		b.uMID50, b.uMID51, b.uMID52, b.uMID53, b.uMID54, b.uMID55, b.uMID56, b.uMID57, b.uMID58, b.uMID59,
		b.uMID60, b.uMID61, b.uMID62, b.uMID63}

	switch b.version {
	case 0:
		reserve := [254]byte{}
		copy(data[348:], reserve[:])
	case 1:
		copy(data[348:412], umidFields[:])
		reserve := [190]byte{}
		copy(data[412:], reserve[:])
	case 2:
		copy(data[348:412], umidFields[:])
		byteOrder.PutUint16(data[412:414], b.loudnessValue)
		byteOrder.PutUint16(data[414:416], b.loudnessRange)
		byteOrder.PutUint16(data[416:418], b.maxTruePeakLevel)
		byteOrder.PutUint16(data[418:420], b.maxMomentaryLoudness)
		byteOrder.PutUint16(data[420:422], b.maxShortTermLoudness)
		reserve := [180]byte{}
		copy(data[422:], reserve[:])
	}

	data = append(data, b.codingHistory...)
	dataSize := len(data)
	b.Header = EncodeChunkHeader(CreateFourCC(BEXTID), uint32(dataSize), byteOrder)
	bytes := append(b.Header.Bytes(), data...)

	return pad(bytes)
}

// DecodeBextChunk provided byte array to Bext.
//
// Array content should be:
// chunk header - 8 bytes (min. requirement for successful decoding)
// data - 602 bytes
// coding history - not restricted amount of bytes
func DecodeBextChunk(data []byte) (*Bext, error) {
	if len(data) < int(HeaderSizeBytes) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", HeaderSizeBytes)
		return nil, errors.New(msg)
	}

	b := &Bext{}
	byteOrder := binary.LittleEndian
	b.Header = decodeChunkHeader(data[:HeaderSizeBytes], 0, byteOrder)
	buf := bytes.NewReader(data[HeaderSizeBytes:])
	fields := []interface{}{&b.description, &b.originator, &b.originatorReference,
		&b.originationDate, &b.originationTime, &b.timeReferenceLow, &b.timeReferenceHigh, &b.version,
		&b.uMID0, &b.uMID1, &b.uMID2, &b.uMID3, &b.uMID4, &b.uMID5, &b.uMID6, &b.uMID7, &b.uMID8, &b.uMID9,
		&b.uMID10, &b.uMID11, &b.uMID12, &b.uMID13, &b.uMID14, &b.uMID15, &b.uMID16, &b.uMID17, &b.uMID18, &b.uMID19,
		&b.uMID20, &b.uMID21, &b.uMID22, &b.uMID23, &b.uMID24, &b.uMID25, &b.uMID26, &b.uMID27, &b.uMID28, &b.uMID29,
		&b.uMID30, &b.uMID31, &b.uMID32, &b.uMID33, &b.uMID34, &b.uMID35, &b.uMID36, &b.uMID37, &b.uMID38, &b.uMID39,
		&b.uMID40, &b.uMID41, &b.uMID42, &b.uMID43, &b.uMID44, &b.uMID45, &b.uMID46, &b.uMID47, &b.uMID48, &b.uMID49,
		&b.uMID50, &b.uMID51, &b.uMID52, &b.uMID53, &b.uMID54, &b.uMID55, &b.uMID56, &b.uMID57, &b.uMID58, &b.uMID59,
		&b.uMID60, &b.uMID61, &b.uMID62, &b.uMID63,
		&b.loudnessValue, &b.loudnessRange, &b.maxTruePeakLevel, &b.maxMomentaryLoudness, &b.maxShortTermLoudness, &b.reserved}

	for _, f := range fields {
		err := binary.Read(buf, byteOrder, f)

		if err != nil {
			err = handleError(err)

			return b, err
		}
	}

	b.codingHistory = make([]byte, buf.Len())
	err := binary.Read(buf, byteOrder, &b.codingHistory)

	if err != nil {
		err = handleError(err)

		return b, err
	}

	return b, nil
}
