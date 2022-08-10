package cipher

import "encoding/base64"

// DynamicAESCEncryptResult RSA + AES 动态加密结果
type DynamicAESCEncryptResult struct {
	AESKey                string
	EncryptedAESKeyBase64 string
	EncryptedData         []byte
	Err                   error
}

// DynamicAESEncryptWithCertificatePEM RSA + AES 动态加密结果（使用 PEM 格式证书）
func DynamicAESEncryptWithCertificatePEM(certPEMData, msg []byte) *DynamicAESCEncryptResult {
	pubKeyPEMData, err := ExtractPublicKeyPEMFromCertificatePEM(certPEMData)
	if err != nil {
		return &DynamicAESCEncryptResult{
			AESKey:                "",
			EncryptedAESKeyBase64: "",
			EncryptedData:         nil,
			Err:                   err,
		}
	}

	return DynamicAESEncryptWithPublicKeyPEM(pubKeyPEMData, msg)
}

// DynamicAESEncryptWithPublicKeyPEM RSA + AES 动态加密结果（使用 PEM 格式公钥）
func DynamicAESEncryptWithPublicKeyPEM(pubKeyPEMData, msg []byte) *DynamicAESCEncryptResult {

	// 生成 16 字节字符串作为 AES 秘钥
	aesKey := GenerateRandomString16()

	// 使用 RSA 为 AES 秘钥加密
	encryptedAESKeyData, err := EncryptWithPublicKey(pubKeyPEMData, []byte(aesKey))
	if err != nil {
		return &DynamicAESCEncryptResult{
			AESKey:                "",
			EncryptedAESKeyBase64: "",
			EncryptedData:         nil,
			Err:                   err,
		}
	}

	// AES 秘钥的 Base64
	encryptedAESKeyBase64 := base64.StdEncoding.EncodeToString(encryptedAESKeyData)

	// 为数据 AES 加密
	encryptedData := AesEncrypt(msg, []byte(aesKey))

	return &DynamicAESCEncryptResult{
		AESKey:                aesKey,
		EncryptedAESKeyBase64: encryptedAESKeyBase64,
		EncryptedData:         encryptedData,
		Err:                   nil,
	}
}
