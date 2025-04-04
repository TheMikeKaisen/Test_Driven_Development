package di

import (
	"io"
)

func Logger(writer io.Writer, name string) error {
	_, err := writer.Write([]byte("Hello, " + name))
	return err
}


