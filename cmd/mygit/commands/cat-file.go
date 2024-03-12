package commands

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CatFile(hash string) (string, error) {
	dirPath, filePath := hash[:2], hash[2:]
	fullPath := dirPath + string(os.PathSeparator) + filePath
	compressedData, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	data, err := Decompress(compressedData)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Decompress(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	decomressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decomressed, nil
}
