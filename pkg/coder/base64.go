package coder

import "encoding/base64"

// Base64StdEncode Base64 标准编码
func Base64StdEncode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64StdDecode Base64 标准解码
func Base64StdDecode(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}
