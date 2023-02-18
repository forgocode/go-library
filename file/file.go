package file

import (
	"io"
	"os"
)

func ReadLargeFile(fileName string, handle func([]byte) error) error {
	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return nil
		}
		if err == io.EOF {
			break
		}
		if handle(buf[:n]) != nil {
			return err
		}
	}
	return nil
}

func ReadFile(fileName string) (string, error) {
	str, err := os.ReadFile(fileName)
	return string(str), err
}

func IsFileExist(fileName string) bool {
	f, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}

func WriteFile(fileName string, msg []byte) {

}
