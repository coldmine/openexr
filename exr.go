package exr

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"os"
)

var MagicNumber = []byte{0x76, 0x2f, 0x31, 0x01}

func Decode(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)
	magic := make([]byte, 4)
	r.Read(magic)
	if !bytes.Equal(magic, MagicNumber) {
		return nil, fmt.Errorf("wrong magic number: %v, need %v", magic, MagicNumber)
	}

	// Version field: 4 bytes
	// first byte: version number
	// 2-4  bytes: set of boolean flags
	versionByte := make([]byte, 4)
	r.Read(versionByte)
	version := int(versionByte[0])
	fmt.Println(version)

	// TODO: parse boolean flags
	return nil, nil
}

func fromScanLineFile() {}

func fromSinglePartFile() {}

func fromMultiPartFile() {}
