package chunk

import (
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
)

// IXMLBext for Broadcast Wave Format (BWF) information.
type IXMLBext struct {
	Description          string `xml:"BWF_DESCRIPTION,omitempty"`
	Originator           string `xml:"BWF_ORIGINATOR,omitempty"`
	OriginatorReference  string `xml:"BWF_ORIGINATOR_REFERENCE,omitempty"`
	OriginationDate      string `xml:"BWF_ORIGINATION_DATE,omitempty"`
	OriginationTime      string `xml:"BWF_ORIGINATION_TIME,omitempty"`
	TimeReferenceLow     string `xml:"BWF_TIME_REFERENCE_LOW,omitempty"`
	TimeReferenceHigh    string `xml:"BWF_TIME_REFERENCE_HIGH,omitempty"`
	Version              string `xml:"BWF_VERSION,omitempty"`
	UMID                 string `xml:"BWF_UMID,omitempty"`
	Reserved             string `xml:"BWF_RESERVED,omitempty"`
	CodingHistory        string `xml:"BWF_CODING_HISTORY,omitempty"`
	LoudnessValue        string `xml:"BWF_LOUDNESS_VALUE,omitempty"`
	LoudnessRange        string `xml:"BWF_LOUDNESS_RANGE,omitempty"`
	MaxTruePeakLevel     string `xml:"BWF_MAX_TRUE_PEAK_LEVEL,omitempty"`
	MaxMomentaryLoudness string `xml:"BWF_MAX_MOMENTARY_LOUDNESS,omitempty"`
	MaxShortTermLoudness string `xml:"BWF_MAX_SHORT_TERM_LOUDNESS,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLBext) String() string {
	return fmt.Sprintf("Description: %s\nOriginator: %s\nOriginatorReference: %s\nOriginationDate: %s\nOriginationTime: %s\nTimeReferenceLow: %s\nTimeReferenceHigh: %s\nVersion: %s\nUMID: %s\nReserved: %s\nCodingHistory: %s\nLoudnessValue: %s\nLoudnessRange: %s\nMaxTruePeakLevel: %s\nMaxMomentaryLoudness: %s\nMaxShortTermLoudness: %s\n",
		c.Description, c.Originator, c.OriginatorReference, c.OriginationDate, c.OriginationTime, c.TimeReferenceLow, c.TimeReferenceHigh, c.Version, c.UMID, c.Reserved, c.CodingHistory, c.LoudnessValue, c.LoudnessRange, c.MaxTruePeakLevel, c.MaxMomentaryLoudness, c.MaxShortTermLoudness)
}

// IXMLSyncPointList
type IXMLSyncPointList struct {
	SyncPointCount string           `xml:"SYNC_POINT_COUNT,omitempty"`
	SyncPoints     []*IXMLSyncPoint `xml:"SYNC_POINT,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLSyncPointList) String() string {
	return fmt.Sprintf("SyncPointCount: %s\n", c.SyncPointCount)
}

// IXMLSyncPoint
type IXMLSyncPoint struct {
	SyncPointType          string `xml:"SYNC_POINT_TYPE,omitempty"`
	SyncPointFunction      string `xml:"SYNC_POINT_FUNCTION,omitempty"`
	SyncPointComment       string `xml:"SYNC_POINT_COMMENT,omitempty"`
	SyncPointLow           string `xml:"SYNC_POINT_LOW,omitempty"`
	SyncPointHigh          string `xml:"SYNC_POINT_HIGH,omitempty"`
	SyncPointEventDuration string `xml:"SYNC_POINT_EVENT_DURATION,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLSyncPoint) String() string {
	return fmt.Sprintf("SyncPointType: %s\nSyncPointFunction: %s\nSyncPointComment: %s\nSyncPointLow: %s\nSyncPointHigh: %s\nSyncPointEventDuration: %s\n",
		c.SyncPointType, c.SyncPointFunction, c.SyncPointComment, c.SyncPointLow, c.SyncPointHigh, c.SyncPointEventDuration)
}

// IXMLSpeed for speed information.
type IXMLSpeed struct {
	Note                            string `xml:"NOTE,omitempty"`
	MasterSpeed                     string `xml:"MASTER_SPEED,omitempty"`
	CurrentSpeed                    string `xml:"CURRENT_SPEED,omitempty"`
	TimecodeRate                    string `xml:"TIMECODE_RATE,omitempty"`
	TimecodeFlag                    string `xml:"TIMECODE_FLAG,omitempty"`
	FileSampleRate                  string `xml:"FILE_SAMPLE_RATE,omitempty"`
	AudioBitDepth                   string `xml:"AUDIO_BIT_DEPTH,omitempty"`
	DigitizerSampleRate             string `xml:"DIGITIZER_SAMPLE_RATE,omitempty"`
	TimestampSamplesSinceMidnightHi string `xml:"TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_HI,omitempty"`
	TimestampSamplesSinceMidnightLo string `xml:"TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO,omitempty"`
	TimestampSampleRate             string `xml:"TIMESTAMP_SAMPLE_RATE,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLSpeed) String() string {
	return fmt.Sprintf("Note: %s\nMasterSpeed: %s\nCurrentSpeed: %s\nTimecodeRate: %s\nTimecodeFlag: %s\nFileSampleRate: %s\nAudioBitDepth: %s\nDigitizerSampleRate: %s\nTimestampSamplesSinceMidnightHi: %s\nTimestampSamplesSinceMidnightLo: %s\nTimestampSampleRate:%s\n",
		c.Note, c.MasterSpeed, c.CurrentSpeed, c.TimecodeRate, c.TimecodeFlag, c.FileSampleRate, c.AudioBitDepth, c.DigitizerSampleRate, c.TimestampSamplesSinceMidnightHi, c.TimestampSamplesSinceMidnightLo, c.TimestampSampleRate)
}

// IXMLLoudness for ludness information equivalent to thos in IXMLBext.
type IXMLLoudness struct {
	LoudnessValue        string `xml:"LOUDNESS_VALUE,omitempty"`
	LoudnessRange        string `xml:"LOUDNESS_RANGE,omitempty"`
	MaxTruePeakLevel     string `xml:"MAX_TRUE_PEAK_LEVEL,omitempty"`
	MaxMomentaryLoudness string `xml:"MAX_MOMENTARY_LOUDNESS,omitempty"`
	MaxShortTermLoudness string `xml:"MAX_SHORT_TERM_LOUDNESS,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLLoudness) String() string {
	return fmt.Sprintf("LoudnessValue: %s\nLoudnessRange: %s\nMaxTruePeakLevel: %s\nMaxMomentaryLoudness: %s\nMaxShortTermLoudness: %s\n",
		c.LoudnessRange, c.LoudnessValue, c.MaxTruePeakLevel, c.MaxMomentaryLoudness, c.MaxShortTermLoudness)
}

// IXMLHistory for tracking a file's origins.
type IXMLHistory struct {
	OriginalFileName string `xml:"ORIGINAL_FILENAME,omitempty"`
	ParentFilename   string `xml:"PARENT_FILERNAME,omitempty"`
	ParentUID        string `xml:"PARENT_UID,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLHistory) String() string {
	return fmt.Sprintf("OriginalFileName: %s\nParentFilename: %s\nParentUID: %s\n",
		c.OriginalFileName, c.ParentFilename, c.ParentUID)
}

// IXMLFileSet information for grouping of recorded files.
type IXMLFileSet struct {
	TotalFiles   string `xml:"TOTAL_FILES,omitempty"`
	FamilyUID    string `xml:"FAMILY_UID,omitempty"`
	FamilyName   string `xml:"FAMILY_NAME,omitempty"`
	FileSetIndex string `xml:"FILE_SET_INDEX,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLFileSet) String() string {
	return fmt.Sprintf("TotalFiles: %s\nFamilyUID: %s\nFamilyName: %s\nFileSetIndex: %s\n",
		c.TotalFiles, c.FamilyUID, c.FamilyName, c.FileSetIndex)
}

// IXMLTrackList for track identification.
type IXMLTrackList struct {
	TrackCount string       `xml:"TRACK_COUNT,omitempty"`
	Tracks     []*IXMLTrack `xml:"TRACK,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLTrackList) String() string {
	return fmt.Sprintf("TrackCount: %s\n", c.TrackCount)
}

// IXMLTrack for track identification.
type IXMLTrack struct {
	ChannelIndex    string `xml:"CHANNEL_INDEX,omitempty"`
	InterleaveIndex string `xml:"INTERLEAVE_INDEX,omitempty"`
	Name            string `xml:"NAME,omitempty"`
	Function        string `xml:"FUNCTION,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLTrack) String() string {
	return fmt.Sprintf("ChannelIndex: %s\nInterleaveIndex: %s\nName: %s\nFunction: %s\n",
		c.ChannelIndex, c.InterleaveIndex, c.Name, c.Function)
}

// IXMLUser for user information.
type IXMLUser struct {
	FullTitle                 string `xml:"FULL_TITLE,omitempty"`
	DirecorName               string `xml:"DIRECTOR_NAME,omitempty"`
	ProductionName            string `xml:"PRODUCTION_NAME,omitempty"`
	ProductionAddress         string `xml:"PRODUCTION_ADDRESS,omitempty"`
	ProductionEmail           string `xml:"PRODUCTION_EMAIL,omitempty"`
	ProductionPhone           string `xml:"PRODUCTION_PHONE,omitempty"`
	ProductionNote            string `xml:"PRODUCTION_NOTE,omitempty"`
	SoundMixerName            string `xml:"SOUND_MIXER_NAME,omitempty"`
	SoundMixerAddress         string `xml:"SOUND_MIXER_ADDRESS,omitempty"`
	SoundMixerEmail           string `xml:"SOUND_MIXER_EMAIL,omitempty"`
	SoundMixerPhone           string `xml:"SOUND_MIXER_PHONE,omitempty"`
	SoundMixerNote            string `xml:"SOUND_MIXER_NOTE,omitempty"`
	AudioRecorderModel        string `xml:"AUDIO_RECORDER_MODEL,omitempty"`
	AudioRecorderSerialNumber string `xml:"AUDIO_RECORDER_SERIAL_NUMBER,omitempty"`
	AudioRecorderFirmware     string `xml:"AUDIO_RECORDER_FIRMWARE,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLUser) String() string {
	return fmt.Sprintf("FullTitle: %s\nDirecorName: %s\nProductionName: %s\nProductionAddress: %s\nProductionEmail: %s\nProductionPhone: %s\nProductionNote: %s\nSoundMixerName: %s\nSoundMixerAddress: %s\nSoundMixerEmail: %s\nSoundMixerPhone: %s\nSoundMixerNote: %s\nAudioRecorderModel: %s\nAudioRecorderSerialNumber: %s\nAudioRecorderFirmware: %s\n",
		c.FullTitle, c.DirecorName, c.ProductionName, c.ProductionAddress, c.ProductionEmail, c.ProductionPhone, c.ProductionNote, c.SoundMixerName, c.SoundMixerAddress, c.SoundMixerEmail, c.SoundMixerPhone, c.SoundMixerNote, c.AudioRecorderModel, c.AudioRecorderSerialNumber, c.AudioRecorderFirmware)
}

// IXMLLocation for location description.
type IXMLLocation struct {
	Name     string `xml:"LOCATION_NAME,omitempty"`
	GPS      string `xml:"LOCATION_GPS,omitempty"`
	Altitude string `xml:"LOCATION_ALTITUDE,omitempty"`
	Type     string `xml:"LOCATION_TYPE,omitempty"`
	Time     string `xml:"LOCATION_TIME,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXMLLocation) String() string {
	return fmt.Sprintf("Name: %s\nGPS: %s\nAltitude: %s\nType: %s\nTime: %s\n",
		c.Name, c.GPS, c.Altitude, c.Type, c.Time)
}

// IXML for providing project based metadata of production as specified in the iXML specification.
type IXML struct {
	*Header
	IXMLVersion          string             `xml:"IXML_VERSION,omitempty"`
	Project              string             `xml:"PROJECT,omitempty"`
	Scene                string             `xml:"SCENE,omitempty"`
	Tape                 string             `xml:"TAPE,omitempty"`
	Take                 string             `xml:"TAKE,omitempty"`
	TakeType             string             `xml:"TAKE_TYPE,omitempty"`
	NoGood               bool               `xml:"NO_GOOD,omitempty"`
	FalseStart           bool               `xml:"FALSE_START,omitempty"`
	WildTrack            bool               `xml:"WILD_TAKE,omitempty"`
	Circled              bool               `xml:"CIRCLED,omitempty"`
	FileUID              string             `xml:"FILE_UID,omitempty"`
	Ubits                string             `xml:"UBITS,omitempty"`
	Note                 string             `xml:"NOTE,omitempty"`
	SyncPointList        *IXMLSyncPointList `xml:"SYNC_POINT_LIST,omitempty"`
	Speed                *IXMLSpeed         `xml:"SPEED,omitempty"`
	Loudness             *IXMLLoudness      `xml:"LOUDNESS,omitempty"`
	History              *IXMLHistory       `xml:"HISTORY,omitempty"`
	FileSet              *IXMLFileSet       `xml:"FILE_SET,omitempty"`
	TrackList            *IXMLTrackList     `xml:"TRACK_LIST,omitempty"`
	PreRecordSamplecount string             `xml:"PRE_RECORD_SAMPLECOUNT,omitempty"`
	Bext                 *IXMLBext          `xml:"BEXT,omitempty"`
	User                 *IXMLUser          `xml:"USER,omitempty"`
	Location             *IXMLLocation      `xml:"LOCATION,omitempty"`
}

// String returns string represensation of chunk.
func (c *IXML) String() string {
	return fmt.Sprintf("IXMLVersion: %s\nProject: %s\nScene: %s\nTape: %s\nTake: %s\nTakeType: %s\nNoGood: %t\nFalseStart: %t\nWildTrack: %t\nCircled: %t\nFileUID: %s\nUBits: %s\nNote: %s\nSyncPointList: %s\nSpeed: %s\nLoudness: %s\nHistory: %s\nFileSet: %s\nTrackList: %s\nPreRecordSampleCount: %s\nBext: %s\nUser: %s\nLocation: %s\n",
		c.IXMLVersion, c.Project, c.Scene, c.Tape, c.Take, c.TakeType, c.NoGood, c.FalseStart, c.WildTrack, c.Circled, c.FileUID, c.Ubits, c.Note, c.SyncPointList, c.Speed, c.Loudness, c.History, c.FileSet, c.TrackList, c.PreRecordSamplecount, c.Bext, c.User, c.Location)
}

// Bytes converts IXML to byte array. A new Header with id 'iXML' is created.
//
// Header size is set to real data size. A minimum amount of 25 bytes is returned.
// chunk header - 8 bytes
// XML root - 17 bytes '<BWFXML></BWFXML>'
//
// A padding byte is added if size is odd. This optional byte is not reflected in size.
func (c *IXML) Bytes() ([]byte, error) {
	tmp := struct {
		*IXML
		XMLName struct{} `xml:"BWFXML"`
	}{IXML: c}

	data, err := xml.Marshal(tmp)

	if err != nil {
		return nil, err
	}

	dataSize := len(data)
	header := EncodeChunkHeader(CreateFourCC(IXMLID), uint32(dataSize), binary.LittleEndian)
	bytes := append(header.Bytes(), data...)

	return pad(bytes), nil
}

// EncodeIXMLChunk returns encoded chunk 'iXML' from provided byte array.
// Byte array is the XML document without chunk Header as specified in the iXML specification.
func EncodeIXMLChunk(data []byte) (*IXML, error) {
	header := EncodeChunkHeader(CreateFourCC(IXMLID), uint32(len(data)), binary.LittleEndian)
	chunk := &IXML{Header: header}
	err := xml.Unmarshal(data[:], chunk)

	return chunk, err
}

// DecodeIXMLChunk decodes provided byte array to IXML.
//
// Array content should be:
// chunk header - 8 bytes (min. requirement for successful decoding)
// XML - minimum 17 bytes of XML data as specified in the iXML specification
func DecodeIXMLChunk(data []byte) (*IXML, error) {
	if len(data) < int(HeaderSizeBytes) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", HeaderSizeBytes)
		return nil, errors.New(msg)
	}

	header := decodeChunkHeader(data[:HeaderSizeBytes], 0, binary.LittleEndian)
	chunk := &IXML{Header: header}
	err := xml.Unmarshal(data[HeaderSizeBytes:], chunk)

	return chunk, err
}
