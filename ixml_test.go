package chunk

import (
	"strconv"
	"testing"
)

func TestIXML(t *testing.T) {
	chunk := &IXML{}
	bytes, err := chunk.Bytes()

	assertNil(t, err, "err when Bytes of IXML")
	assertEqual(t, len(bytes), 26, "bytes length after DecodeIXMLChunk")

	ixml, err := DecodeIXMLChunk(bytes)
	assertEqual(t, ixml.ID(), IXMLID, "id after DecodeIXMLChunk")
	assertEqual(t, ixml.Size(), 25-HeaderSizeBytes, "size after DecodeIXMLChunk")
	assertEqual(t, ixml.HasPadding(), true, "padding after DecodeIXMLChunk")

	chunk.IXMLVersion = "version"
	chunk.Project = "project"
	chunk.Scene = "scene"
	chunk.Tape = "tape"
	chunk.TakeType = "take type"
	chunk.NoGood = true
	chunk.FalseStart = true
	chunk.WildTrack = true
	chunk.Circled = true
	chunk.FileUID = "FileUID0000"
	chunk.Ubits = "0000000000000111111111111111"
	chunk.Note = "note"
	chunk.PreRecordSamplecount = "1000"

	bytes, err = chunk.Bytes()
	assertNil(t, err, "err")
	assertNotNil(t, bytes, "bytes")
	ixml, err = DecodeIXMLChunk(bytes)
	assertNil(t, err, "err")
	assertNotNil(t, ixml, "iXML")
	assertEqual(t, ixml.IXMLVersion, chunk.IXMLVersion, "iXMLVersion")
	assertEqual(t, ixml.Project, chunk.Project, "project")
	assertEqual(t, ixml.Scene, chunk.Scene, "scene")
	assertEqual(t, ixml.Tape, chunk.Tape, "tape")
	assertEqual(t, ixml.TakeType, chunk.TakeType, "take type")
	assertEqual(t, ixml.NoGood, true, "scene")
	assertEqual(t, ixml.FalseStart, true, "scene")
	assertEqual(t, ixml.WildTrack, true, "scene")
	assertEqual(t, ixml.Circled, true, "scene")
	assertEqual(t, ixml.FileUID, chunk.FileUID, "fileUUID")
	assertEqual(t, ixml.Ubits, chunk.Ubits, "ubits")
	assertEqual(t, ixml.PreRecordSamplecount, chunk.PreRecordSamplecount, "preRecordSamplecount")
	assertEqual(t, ixml.Note, chunk.Note, "note")
	assertNil(t, ixml.SyncPointList, "syncPointList")
	assertNil(t, ixml.Speed, "speed")
	assertNil(t, ixml.Loudness, "loudness")
	assertNil(t, ixml.History, "history")
	assertNil(t, ixml.FileSet, "fileSet")
	assertNil(t, ixml.TrackList, "trackList")
	assertNil(t, ixml.Bext, "bext")
	assertNil(t, ixml.User, "user")
	assertNil(t, ixml.Location, "location")

	chunk.Bext = &IXMLBext{}
	chunk.Bext.Description = "description"

	chunk.Bext.Originator = "originator"
	chunk.Bext.OriginatorReference = "originator reference"
	chunk.Bext.OriginationDate = "originattion date"
	chunk.Bext.OriginationTime = "origination time"
	chunk.Bext.TimeReferenceLow = "timeReferenceLow"
	chunk.Bext.TimeReferenceHigh = "timeReferenceHigh"
	chunk.Bext.Version = "version"
	chunk.Bext.UMID = "umid"
	chunk.Bext.Reserved = "reserved"
	chunk.Bext.CodingHistory = "codingHistory"
	chunk.Bext.LoudnessValue = "loudnessValue"
	chunk.Bext.LoudnessRange = "loudnessRange"
	chunk.Bext.MaxTruePeakLevel = "origination maxTruePekalevel"
	chunk.Bext.MaxMomentaryLoudness = "maxMomentaryLoudness"
	chunk.Bext.MaxShortTermLoudness = "maxShortTerminLoudness"

	chunk.SyncPointList = &IXMLSyncPointList{}
	syncPointCount := 2
	chunk.SyncPointList.SyncPointCount = strconv.Itoa(syncPointCount)

	for i := 0; i < syncPointCount; i++ {
		sp := &IXMLSyncPoint{}
		sp.SyncPointType = "type"
		sp.SyncPointFunction = "function"
		sp.SyncPointComment = "comment"
		sp.SyncPointLow = "low"
		sp.SyncPointHigh = "high"
		sp.SyncPointEventDuration = "duration"
		chunk.SyncPointList.SyncPoints = append(chunk.SyncPointList.SyncPoints, sp)
	}

	chunk.Speed = &IXMLSpeed{}
	chunk.Speed.Note = "note"
	chunk.Speed.MasterSpeed = "masterSpeed"
	chunk.Speed.CurrentSpeed = "currentSpeed"
	chunk.Speed.TimecodeRate = "rate"
	chunk.Speed.TimecodeFlag = "flag"
	chunk.Speed.FileSampleRate = "fileSampleRate"
	chunk.Speed.AudioBitDepth = "audioBitDepth"
	chunk.Speed.DigitizerSampleRate = "digitizerSampleRate"
	chunk.Speed.TimestampSamplesSinceMidnightHi = "timestampSamplesSinceMidnightHi"
	chunk.Speed.TimestampSamplesSinceMidnightLo = "timestampSamplesSinceMidnightLo"
	chunk.Speed.TimestampSampleRate = "timestampSampleRate"

	chunk.Loudness = &IXMLLoudness{}
	chunk.Loudness.LoudnessValue = "loudnessValue"
	chunk.Loudness.LoudnessRange = "loudnessRange"
	chunk.Loudness.MaxTruePeakLevel = "maxTruePeakLevel"
	chunk.Loudness.MaxMomentaryLoudness = "maxMomentaryLoudness"
	chunk.Loudness.MaxShortTermLoudness = "maxShortTermLoudness"

	chunk.History = &IXMLHistory{}
	chunk.History.OriginalFileName = "original"
	chunk.History.ParentFilename = "fileName"
	chunk.History.ParentUID = "parentUID"

	chunk.FileSet = &IXMLFileSet{}
	chunk.FileSet.TotalFiles = "totalFiles"
	chunk.FileSet.FamilyUID = "familyUID"
	chunk.FileSet.FamilyName = "familyName"
	chunk.FileSet.FileSetIndex = "fileSetIndex"

	chunk.TrackList = &IXMLTrackList{}
	trackCount := 2
	chunk.TrackList.TrackCount = strconv.Itoa(trackCount)

	for i := 0; i < trackCount; i++ {
		t := &IXMLTrack{}
		t.ChannelIndex = "channelIndex"
		t.InterleaveIndex = "interleaveIndex"
		t.Name = "name"
		t.Function = "function"
		chunk.TrackList.Tracks = append(chunk.TrackList.Tracks, t)
	}

	chunk.User = &IXMLUser{}
	chunk.User.FullTitle = "FullTitle"
	chunk.User.DirecorName = "DirecorName"
	chunk.User.ProductionName = "ProductionName"
	chunk.User.ProductionAddress = "ProductionAddress"
	chunk.User.ProductionEmail = "ProductionEmail"
	chunk.User.ProductionPhone = "ProductionPhone"
	chunk.User.ProductionNote = "ProductionNote"
	chunk.User.SoundMixerName = "SoundMixerName"
	chunk.User.SoundMixerAddress = "SoundMixerAddress"
	chunk.User.SoundMixerEmail = "SoundMixerEmail"
	chunk.User.SoundMixerPhone = "SoundMixerPhone"
	chunk.User.SoundMixerNote = "SoundMixerNote"
	chunk.User.AudioRecorderModel = "AudioRecorderModel"
	chunk.User.AudioRecorderSerialNumber = "AudioRecorderSerialNumber"
	chunk.User.AudioRecorderFirmware = "AudioRecorderFirmware"

	chunk.Location = &IXMLLocation{}
	chunk.Location.Name = "name"
	chunk.Location.GPS = "gps"
	chunk.Location.Altitude = "altitude"
	chunk.Location.Type = "type"
	chunk.Location.Time = "time"

	bytes, err = chunk.Bytes()
	assertNil(t, err, "err")

	ixml, _ = DecodeIXMLChunk(bytes)

	expectedSize := len(bytes) - int(HeaderSizeBytes)

	if ixml.HasPadding() {
		expectedSize -= 1
	}

	assertEqual(t, ixml.Size(), uint32(expectedSize), "size after DecodeIXMLChunk")

	assertEqual(t, ixml.Bext.Description, chunk.Bext.Description, "bext description")
	assertEqual(t, ixml.Bext.Originator, chunk.Bext.Originator, "bext originator")
	assertEqual(t, ixml.Bext.OriginatorReference, chunk.Bext.OriginatorReference, "bext orginatorReference")
	assertEqual(t, ixml.Bext.OriginationDate, chunk.Bext.OriginationDate, "bext originationDate")
	assertEqual(t, ixml.Bext.OriginationTime, chunk.Bext.OriginationTime, "bext originationTime")
	assertEqual(t, ixml.Bext.TimeReferenceLow, chunk.Bext.TimeReferenceLow, "bext timeReferenceLow")
	assertEqual(t, ixml.Bext.TimeReferenceHigh, chunk.Bext.TimeReferenceHigh, "bext timeReferenceHigh")
	assertEqual(t, ixml.Bext.Version, chunk.Bext.Version, "bext version")
	assertEqual(t, ixml.Bext.UMID, chunk.Bext.UMID, "bext umid")
	assertEqual(t, ixml.Bext.Reserved, chunk.Bext.Reserved, "bext reserved")
	assertEqual(t, ixml.Bext.CodingHistory, chunk.Bext.CodingHistory, "bext codingHistory")
	assertEqual(t, ixml.Bext.LoudnessValue, chunk.Bext.LoudnessValue, "bext loudnessValue")
	assertEqual(t, ixml.Bext.LoudnessRange, chunk.Bext.LoudnessRange, "bext loudnessRange")
	assertEqual(t, ixml.Bext.MaxTruePeakLevel, chunk.Bext.MaxTruePeakLevel, "bext maxTruePeakLevel")
	assertEqual(t, ixml.Bext.MaxMomentaryLoudness, chunk.Bext.MaxMomentaryLoudness, "bext maxMomentaryLoudness")
	assertEqual(t, ixml.Bext.MaxShortTermLoudness, chunk.Bext.MaxShortTermLoudness, "bext maxShortTermLoudness")

	assertEqual(t, ixml.SyncPointList.SyncPointCount, chunk.SyncPointList.SyncPointCount, "syncPointCount")
	assertEqual(t, len(ixml.SyncPointList.SyncPoints), len(chunk.SyncPointList.SyncPoints), "syncPoints")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointType, chunk.SyncPointList.SyncPoints[0].SyncPointType, "SyncPointType")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointFunction, chunk.SyncPointList.SyncPoints[0].SyncPointFunction, "SyncPointFunction")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointComment, chunk.SyncPointList.SyncPoints[0].SyncPointComment, "SyncPointComment")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointLow, chunk.SyncPointList.SyncPoints[0].SyncPointLow, "SyncPointLow")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointHigh, chunk.SyncPointList.SyncPoints[0].SyncPointHigh, "SyncPointHigh")
	assertEqual(t, ixml.SyncPointList.SyncPoints[0].SyncPointEventDuration, chunk.SyncPointList.SyncPoints[0].SyncPointEventDuration, "SyncPointEventDuration")

	assertEqual(t, ixml.Speed.Note, chunk.Speed.Note, "Note")
	assertEqual(t, ixml.Speed.MasterSpeed, chunk.Speed.MasterSpeed, "MasterSpeed")
	assertEqual(t, ixml.Speed.CurrentSpeed, chunk.Speed.CurrentSpeed, "CurrentSpeed")
	assertEqual(t, ixml.Speed.TimecodeRate, chunk.Speed.TimecodeRate, "TimecodeRate")
	assertEqual(t, ixml.Speed.TimecodeFlag, chunk.Speed.TimecodeFlag, "TimecodeFlag")
	assertEqual(t, ixml.Speed.FileSampleRate, chunk.Speed.FileSampleRate, "FileSampleRate")
	assertEqual(t, ixml.Speed.AudioBitDepth, chunk.Speed.AudioBitDepth, "AudioBitDepth")
	assertEqual(t, ixml.Speed.DigitizerSampleRate, chunk.Speed.DigitizerSampleRate, "DigitizerSampleRate")
	assertEqual(t, ixml.Speed.TimestampSamplesSinceMidnightHi, chunk.Speed.TimestampSamplesSinceMidnightHi, "TimestampSamplesSinceMidnightHi")
	assertEqual(t, ixml.Speed.TimestampSamplesSinceMidnightLo, chunk.Speed.TimestampSamplesSinceMidnightLo, "TimestampSamplesSinceMidnightLo")
	assertEqual(t, ixml.Speed.TimestampSampleRate, chunk.Speed.TimestampSampleRate, "TimestampSampleRate")

	assertEqual(t, ixml.Loudness.LoudnessValue, chunk.Loudness.LoudnessValue, "LoudnessValue")
	assertEqual(t, ixml.Loudness.LoudnessRange, chunk.Loudness.LoudnessRange, "LoudnessRange")
	assertEqual(t, ixml.Loudness.MaxTruePeakLevel, chunk.Loudness.MaxTruePeakLevel, "MaxTruePeakLevel")
	assertEqual(t, ixml.Loudness.MaxMomentaryLoudness, chunk.Loudness.MaxMomentaryLoudness, "MaxMomentaryLoudness")
	assertEqual(t, ixml.Loudness.MaxShortTermLoudness, chunk.Loudness.MaxShortTermLoudness, "MaxShortTermLoudness")

	assertEqual(t, ixml.History.OriginalFileName, chunk.History.OriginalFileName, "OriginalFileName")
	assertEqual(t, ixml.History.ParentFilename, chunk.History.ParentFilename, "ParentFilename")
	assertEqual(t, ixml.History.ParentUID, chunk.History.ParentUID, "ParentUID")

	assertEqual(t, ixml.FileSet.TotalFiles, chunk.FileSet.TotalFiles, "TotalFiles")
	assertEqual(t, ixml.FileSet.FamilyUID, chunk.FileSet.FamilyUID, "FamilyUID")
	assertEqual(t, ixml.FileSet.FamilyName, chunk.FileSet.FamilyName, "FamilyName")
	assertEqual(t, ixml.FileSet.FileSetIndex, chunk.FileSet.FileSetIndex, "FileSetIndex")

	assertEqual(t, ixml.TrackList.TrackCount, chunk.TrackList.TrackCount, "trackCount")
	assertEqual(t, len(ixml.TrackList.Tracks), len(chunk.TrackList.Tracks), "tracks")
	assertEqual(t, ixml.TrackList.Tracks[0].ChannelIndex, chunk.TrackList.Tracks[0].ChannelIndex, "ChannelIndex")
	assertEqual(t, ixml.TrackList.Tracks[0].InterleaveIndex, chunk.TrackList.Tracks[0].InterleaveIndex, "InterleaveIndex")
	assertEqual(t, ixml.TrackList.Tracks[0].Name, chunk.TrackList.Tracks[0].Name, "Name")
	assertEqual(t, ixml.TrackList.Tracks[0].Function, chunk.TrackList.Tracks[0].Function, "Function")

	assertEqual(t, ixml.User.FullTitle, chunk.User.FullTitle, "FullTitle")
	assertEqual(t, ixml.User.DirecorName, chunk.User.DirecorName, "DirecorName")
	assertEqual(t, ixml.User.ProductionName, chunk.User.ProductionName, "ProductionName")
	assertEqual(t, ixml.User.ProductionName, chunk.User.ProductionName, "ProductionName")
	assertEqual(t, ixml.User.ProductionEmail, chunk.User.ProductionEmail, "ProductionEmail")
	assertEqual(t, ixml.User.ProductionPhone, chunk.User.ProductionPhone, "ProductionPhone")
	assertEqual(t, ixml.User.ProductionNote, chunk.User.ProductionNote, "ProductionNote")
	assertEqual(t, ixml.User.SoundMixerName, chunk.User.SoundMixerName, "SoundMixerName")
	assertEqual(t, ixml.User.SoundMixerAddress, chunk.User.SoundMixerAddress, "SoundMixerAddress")
	assertEqual(t, ixml.User.SoundMixerEmail, chunk.User.SoundMixerEmail, "SoundMixerEmail")
	assertEqual(t, ixml.User.SoundMixerPhone, chunk.User.SoundMixerPhone, "SoundMixerPhone")
	assertEqual(t, ixml.User.SoundMixerNote, chunk.User.SoundMixerNote, "SoundMixerNote")
	assertEqual(t, ixml.User.AudioRecorderModel, chunk.User.AudioRecorderModel, "AudioRecorderModel")
	assertEqual(t, ixml.User.AudioRecorderSerialNumber, chunk.User.AudioRecorderSerialNumber, "AudioRecorderSerialNumber")
	assertEqual(t, ixml.User.AudioRecorderFirmware, chunk.User.AudioRecorderFirmware, "AudioRecorderFirmware")

	assertEqual(t, ixml.Location.Name, chunk.Location.Name, "Name")
	assertEqual(t, ixml.Location.GPS, chunk.Location.GPS, "GPS")
	assertEqual(t, ixml.Location.Altitude, chunk.Location.Altitude, "Altitude")
	assertEqual(t, ixml.Location.Type, chunk.Location.Type, "Type")
	assertEqual(t, ixml.Location.Time, chunk.Location.Time, "Time")

	ixml, err = EncodeIXMLChunk(bytes[HeaderSizeBytes:])
	assertEqual(t, ixml.Header.Size(), uint32(len(bytes)-int(HeaderSizeBytes)), "size after EncodeIXMLChunk")
	assertEqual(t, ixml.Location.Name, chunk.Location.Name, "Name")
	assertEqual(t, ixml.Location.GPS, chunk.Location.GPS, "GPS")
	assertEqual(t, ixml.Location.Altitude, chunk.Location.Altitude, "Altitude")
	assertEqual(t, ixml.Location.Type, chunk.Location.Type, "Type")
	assertEqual(t, ixml.Location.Time, chunk.Location.Time, "Time")
}

func TestDecodeIXMLChunk(t *testing.T) {
	ixml, err := DecodeIXMLChunk(nil)

	assertNotNil(t, err, "err is nil when DecodeIXMLChunk")
	assertNil(t, ixml, "ixml is not nil when DecodeIXMLChunk")

	data1 := [HeaderSizeBytes - 1]byte{}
	ixml, err = DecodeIXMLChunk(data1[:])

	assertNotNil(t, err, "err is nil when DecodeIXMLChunk")
	assertNil(t, ixml, "ixml is not nil when DecodeIXMLChunk")

	data2 := [HeaderSizeBytes]byte{}
	ixml, err = DecodeIXMLChunk(data2[:])

	assertNotNil(t, err, "err is nil when DecodeIXMLChunk")
	assertNotNil(t, ixml, "ixml is nil when DecodeIXMLChunk")
}
