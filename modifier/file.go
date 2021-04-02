package modifier

// // todo: chunk ids might not be unique. Better delete by offset => GetHeaderByStartPos
// // DeleteChunk
// func (rf *File) DeleteChunk(headerID string, reader io.ReaderAt, writer io.WriterAt) (uint32, error) {
// 	foundHeaders := rf.FindHeaders(headerID)

// 	if foundHeaders == nil {
// 		msg := fmt.Sprintf("chunk not found: %s", headerID)
// 		return 0, errors.New(msg)
// 	}

// 	//todo: find by offset
// 	header := foundHeaders[0]
// 	headers := rf.Headers
// 	sort.Sort(chunk.SortHeadersByStartPosAsc(headers))
// 	writeOffset := header.StartPos()

// 	for i := 0; i < len(headers); i++ {
// 		if headers[i].StartPos() > header.StartPos() {
// 			sectionReader := io.NewSectionReader(reader, int64(headers[i].StartPos()), int64(headers[i].FullSize()))
// 			n, err := moveChunk(writeOffset, headers[i].FullSize(), sectionReader, writer)

// 			if err != nil {
// 				return 0, err
// 			}

// 			writeOffset += uint32(n)
// 		}
// 	}

// 	riffSize := rf.Header.Size() - header.FullSize()
// 	err := rf.UpdateSize(riffSize, writer)

// 	if err != nil {
// 		return 0, err
// 	}

// 	fileSize := rf.Header.FullSize() - header.FullSize()

// 	return fileSize, nil
// }

// func moveChunk(writeOffset uint32, size uint32, reader io.Reader, writer io.WriterAt) (int, error) {
// 	bytes := make([]byte, size)
// 	n, err := io.ReadFull(reader, bytes[:])

// 	if err != nil {
// 		return 0, err
// 	}

// 	_, err = writer.WriteAt(bytes, int64(writeOffset))

// 	if err != nil {
// 		return 0, err
// 	}

// 	return n, nil
// }

// // AddChunk reads chnunk bytes from io.Reader and writes to io.Writer
// func (rf *File) AddChunk(reader io.Reader, writer io.WriterAt, bufferSize int) error {
// 	if bufferSize <= 0 {
// 		bufferSize = 1024
// 	}

// 	// start writing chunk to end of file
// 	offset := int64(rf.Header.FullSize())
// 	var chunkSize uint32 = 0

// 	for {
// 		b := make([]byte, bufferSize)
// 		n, err := reader.Read(b)

// 		if n > 0 {
// 			b = b[:n]
// 			_, err := writer.WriteAt(b, offset)

// 			if err != nil {
// 				return err
// 			}

// 			offset += int64(n)
// 			chunkSize += uint32(n)
// 		}

// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			return err
// 		}
// 	}

// 	riffSize := rf.Header.size + chunkSize
// 	err := rf.UpdateSize(riffSize, writer)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateSize
// func (rf *File) UpdateSize(size uint32, writer io.WriterAt) error {
// 	buf := new(bytes.Buffer)
// 	err := binary.Write(buf, binary.LittleEndian, size)

// 	if err != nil {

// 		return err
// 	}

// 	b := buf.Bytes()
// 	_, err = writer.WriteAt(b, 4)

// 	if err != nil {
// 		return err
// 	}

// 	rf.Header.size = size

// 	return nil
// }
