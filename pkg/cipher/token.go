package cipher

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// GenerateRandomString16 生成 16 字节随机字符串
func GenerateRandomString16() string {
	uuidVal := uuid.New()
	retStr := fmt.Sprintf("%X", md5.Sum(uuidVal[:]))
	retStr = strings.ToLower(retStr)
	retStr = retStr[8:24]
	return retStr
}
