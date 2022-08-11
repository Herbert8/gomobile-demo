package coder

import "net/url"

// URLQueryEscape 针对参数的 URL 编码
func URLQueryEscape(s string) string {
	return url.QueryEscape(s)
}

// URLQueryUnescape 针对参数的 URL 解码
func URLQueryUnescape(s string) (string, error) {
	return url.QueryUnescape(s)
}

// URLPathEscape 针对路径的 URL 编码
func URLPathEscape(s string) string {
	return url.PathEscape(s)
}

// URLPathUnescape 针对路径的 URL 解码
func URLPathUnescape(s string) (string, error) {
	return url.PathUnescape(s)
}
