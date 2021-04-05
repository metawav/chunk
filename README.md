# wav [![GoDoc](https://godoc.org/github.com/pmoule/wav?status.svg)](https://godoc.org/github.com/pmoule/wav)
Library for reading chunked media container files like RIFF/ AIFF.

## Features
- Display headers of contained chunks of RIFF / AIFF / AIFF-C
- Decode / Encode chunks 'fmt ' and 'COMM'
- Decode / Encode chunk header
- Decode / Encode container header

## Prerequisites
Download and install `wav` into your GOPATH.
```
go get github.com/pmoule/wav
```
Import the package to get started.
```go
import "github.com/pmoule/wav"
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
### Find format chunk
```go
headers := container.FindHeaders("fmt ") // kepp care of trailing space :)
chunkBytes := make([]byte, headers[0].FullSize())
_, err = file.ReadAt(chunkBytes, int64(headers[0].StartPos()))
format, _ := chunk.DecodeFMTChunk(chunkBytes)
fmt.Printf("%s\n", format)
```
Example output:
```
Format: 1
Channels: 1
Sample rate: 44100
Byte rate: 88200
Bytes per sample: 2
```
## Documentation
See package documentation:

[![GoDoc](https://godoc.org/github.com/pmoule/wav?status.svg)](https://godoc.org/github.com/pmoule/wav)

## License
`wav` is released under Apache License, Version 2.0. See [LICENSE](LICENSE.txt).