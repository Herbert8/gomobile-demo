package detector

import (
	"bytes"
	"net/http"
)

type sniffSignature interface {
	// match returns the MIME type of the data, or "" if unknown.
	match(data []byte, firstNonWS int) string
}

type exactSignature struct {
	sig []byte
	ct  string
}

func (e *exactSignature) match(data []byte, firstNonWS int) string {
	if bytes.HasPrefix(data, e.sig) {
		return e.ct
	}
	return ""
}

var sniffSignatures = []sniffSignature{
	// 参考 常见 MIME 类型列表
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// 参考 List of file signatures
	// https://en.wikipedia.org/wiki/List_of_file_signatures
	&exactSignature{[]byte("\x37\x7A\xBC\xAF\x27\x1C"), "application/x-7z-compressed"},
}

// DetectContentType 监测数据的 MIME 类型
func DetectContentType(data []byte) string {
	const DefaultContentType = "application/octet-stream"
	contentType := http.DetectContentType(data)
	if contentType != DefaultContentType {
		return contentType
	}
	for _, sig := range sniffSignatures {
		if ct := sig.match(data, 0); ct != "" {
			return ct
		}
	}
	return DefaultContentType
}
