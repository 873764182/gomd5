package gomd5

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

// FileMd5 get file md5
func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("md5.go hash.FileMd5 os open error %v", err))
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy error %v", err))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// StringMd5 get string md5
func StringMd5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
