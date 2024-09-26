package main

import (
	"io"
	"strings"
)

// IO Readerインターフェースでnet.connでもos.fileでも同じ関数で処理できる
func readAndToUpperCase(r io.Reader) (string, error) {
	data := make([]byte, 128)
	c, err := r.Read(data)
	if err != nil {
		return "", err
	}
	upperStr := strings.ToUpper(string(data[:c]))
	return upperStr, nil
}
