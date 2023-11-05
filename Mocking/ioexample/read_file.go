package ioexample

import (
	"io"
)

// ReadFileContent reads the first numBytes bytes from the file
// *os.File é um tipo concreto que possui métodos Read e Close

// func ReadFileContent(f *os.File, numBytes int) ([]byte, error) {
// 	data := make([]byte, numBytes)
// 	_, err := f.Read(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	return data, nil
// }

// Uma forma é usar uma interface que o tipo concreto implementa
func ReadFileContent(rc io.ReadCloser, numBytes int) ([]byte, error) {
	data := make([]byte, numBytes)
	_, err := rc.Read(data)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return data, nil
}

// Também é possível definir a própria interface...?
