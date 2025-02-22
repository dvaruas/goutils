package miscutils

import (
	"fmt"
	"os"
)

type Writer interface {
	WriteLine(format string, a ...any) error
	Close() error
}

type FileWriter struct {
	fw *os.File
}

func NewFileWriter(filePath string) (*FileWriter, error) {
	fw, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return &FileWriter{
		fw: fw,
	}, nil
}

func (fww *FileWriter) WriteLine(format string, a ...any) error {
	s := fmt.Sprintf(format, a...) + "\n"
	_, err := fww.fw.WriteString(s)
	return err
}

func (fww *FileWriter) Close() error {
	return fww.fw.Close()
}

type StdoutWriter struct{}

func (sw *StdoutWriter) WriteLine(format string, a ...any) error {
	_, err := fmt.Printf(format+"\n", a...)
	return err
}

func (sw *StdoutWriter) Close() error {
	return nil
}
