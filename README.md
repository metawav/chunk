# wav
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
## Documentation
See package documentation:

[![GoDoc](https://godoc.org/github.com/pmoule/wav?status.svg)](https://godoc.org/github.com/pmoule/wav)

## License
`wav` is released under Apache License, Version 2.0. See [LICENSE](LICENSE.txt).