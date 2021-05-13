# chunk 
![Build](https://github.com/metawav/chunk//actions/workflows/go_build_and_test.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/metawav/chunk?status.svg)](https://godoc.org/github.com/metawav/chunk)

Library for decoding and encoding chunked audio containers like RIFF / AIFF.

## Features
- Supported formats: 
  - Waveform Audio File Format (WAVE)
  - Broadcast Wave Format (BWF)
  - Audio Interchange File Format (AIFF, AIFF-C)
- Supported chunks (Decode / Encode)
  - 'fmt ' - WAVE format and PCM format
  - 'COMM' - Common
  - 'bext' - Broadcast Extension for sound metadata in BWF (version 0, 1, 2)
  - 'iXML' - Extension for sound metadata in BWF (iXML Specification Revision 2.10)
- Decode headers of unknown chunks
- Encode custom header
- Byte array presentation of known chunks for further processing

## Prerequisites
Download and install `chunk` into your GOPATH.
```
go get github.com/metawav/chunk
```
Import the package to get started.
```go
import "github.com/metawav/chunk"
```
## Examples
### Read headers from RIFF (.wav) file
```go
fileName := "sample.wav"
file, err := os.Open(fileName)

if err != nil {
    panic(err)
}

defer file.Close()

reader := io.ReadSeeker(file)
container, _ := chunk.ReadRiff(file.Name(), reader)

for _, header := range container.Headers {
    fmt.Printf("%s\n", header)
}
```
Example output:
```
ID: fmt  Size: 16 FullSize: 24 StartPos: 12 HasPadding: false
ID: data Size: 88200 FullSize: 88208 StartPos: 36 HasPadding: false
ID: LIST Size: 68 FullSize: 76 StartPos: 88244 HasPadding: false
ID: JUNK Size: 10 FullSize: 18 StartPos: 88320 HasPadding: false
```
### Find and decode format chunk
```go
headers := container.FindHeaders("fmt ") // keep care of trailing space :)
chunkBytes := make([]byte, headers[0].FullSize())
_, err = file.ReadAt(chunkBytes, int64(headers[0].StartPos()))

if err != nil {
    //handle err
}

format, err := chunk.DecodePCMFormatChunk(chunkBytes)

if err != nil {
    //handle err
}

fmt.Printf("%s\n", format)
```
Example output:
```
Format: 1
Channels: 1
Sample rate: 44100
Byte rate: 88200
Bytes per sample: 2
Bits per sample: 16
```
### Encode format chunk
```go
format = chunk.EncodePCMFormatChunk(16, 1, 1, 44100, 88200, 2, 24)
fmt.Printf("%s\n", format)
```
Example output:
```
Format: 1
Channels: 1
Sample rate: 44100
Byte rate: 88200
Bytes per sample: 2
Bits per sample: 24
```
### Get bytes representation of format chunk
```go
format = chunk.EncodePCMFormatChunk(16, 1, 1, 44100, 88200, 2, 24)
fmt.Printf("%08b\n", format.Bytes())
```
Example output:
```
[
01100110 01101101 01110100 00100000 // ID               = "fmt "
00010000 00000000 00000000 00000000 // size             = 16
00000001 00000000                   // format           = 1
00000001 00000000                   // channels         = 1
01000100 10101100 00000000 00000000 // samples per sec  = 44100
10001000 01011000 00000001 00000000 // bytes per sec    = 88200
00000010 00000000                   // block align      = 2
00011000 00000000                   // bits per sample  = 24
]
```
## Documentation
See package documentation:

[![GoDoc](https://godoc.org/github.com/metawav/chunk?status.svg)](https://godoc.org/github.com/metawav/chunk)

## License
`chunk` is released under Apache License, Version 2.0. See [LICENSE](LICENSE.txt).
