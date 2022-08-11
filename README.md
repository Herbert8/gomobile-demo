# Go Mobile Common Utilities



## 一、API 说明

### cert 包

证书链校验及证书公钥提取。

```go
// CheckSignatureFrom 判断当前证书是否继承自另一个证书
func (receiver *CertificateWrapper) CheckSignatureFrom(parentCertWrapperObj *CertificateWrapper) error

// CommonName 获取证书的 CommonName
func (receiver *CertificateWrapper) CommonName() string

// PublicKeyPEMString 获取证书的 PublicKey，格式为 PEM
func (receiver *CertificateWrapper) PublicKeyPEMString() (string, error)



// NewCertificateWrapperFromPEMData 通过 PEM 格式的数据创建证书
func NewCertificateWrapperFromPEMData(pemData []byte) (*CertificateWrapper, error)

// NewCertificateWrapperFromPEMString 通过 PEM 格式的字符串创建证书
func NewCertificateWrapperFromPEMString(pemStr string) (*CertificateWrapper, error)

// CheckPEMCertSignatureFromParentPEMCertString 通过 PEM 格式字符串判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertString(pemCertStr string, parentPemCertStr string) error

// CheckPEMCertSignatureFromParentPEMCertData 通过 PEM 格式数据判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertData(pemCertData []byte, parentPemCertData []byte) error
```

### archive 包

#### 7Zip

7Zip 读取器。

```go
// FilesIn7ZipArchive 获取压缩包内的文件列表
func FilesIn7ZipArchive(sevenZipFile string, password string) (string, error)
// Extract7ZipArchive 解压缩
func Extract7ZipArchive(sevenZipFile string, password string, rootPath string) error
```



### cipher 包

加密相关工具。

#### AES

```go
// SetIVString 设置 IV
func SetIVString(ivStr string)
// GetIVString 获取 IV
func GetIVString() string
// AesEncrypt AES加密
func AesEncrypt(originalData, key []byte) []byte
// AesDecrypt AES解密
func AesDecrypt(secretData, key []byte) []byte
```

#### RSA

```go
// ExtractPublicKeyPEMFromCertificatePEM 获取证书的 PublicKey，格式为 PEM
func ExtractPublicKeyPEMFromCertificatePEM(certPEMData []byte) ([]byte, error)
// EncryptWithCertificate 使用证书对数据进行加密
func EncryptWithCertificate(certPEMData []byte, msg []byte) ([]byte, error) 
// EncryptWithPublicKey 使用公钥对数据进行加密
func EncryptWithPublicKey(pubKeyPEMData []byte, msg []byte) ([]byte, error)
```

#### 混合加密

```go
// DynamicAESCEncryptResult RSA + AES 动态加密结果
type DynamicAESCEncryptResult struct {
	AESKey                string
	EncryptedAESKeyBase64 string
	EncryptedData         []byte
	Err                   error
}
// DynamicAESEncryptWithCertificatePEM RSA + AES 动态加密结果（使用 PEM 格式证书）
func DynamicAESEncryptWithCertificatePEM(certPEMData, msg []byte) *DynamicAESCEncryptResult
// DynamicAESEncryptWithPublicKeyPEM RSA + AES 动态加密结果（使用 PEM 格式公钥）
func DynamicAESEncryptWithPublicKeyPEM(pubKeyPEMData, msg []byte) *DynamicAESCEncryptResult 
```

#### token

```go
// GenerateRandomString16 生成 16 字节随机字符串
func GenerateRandomString16() string
```



### coder 包

常用编解码。

#### Base64

```go
// Base64StdEncode Base64 标准编码
func Base64StdEncode(data []byte) string
// Base64StdDecode Base64 标准解码
func Base64StdDecode(base64Str string) ([]byte, error)
```

#### URL

```go
// URLQueryEscape 针对参数的 URL 编码
func URLQueryEscape(s string) string
// URLQueryUnescape 针对参数的 URL 解码
func URLQueryUnescape(s string) (string, error)
// URLPathEscape 针对路径的 URL 编码
func URLPathEscape(s string) string
// URLPathUnescape 针对路径的 URL 解码
func URLPathUnescape(s string) (string, error)
```



### detector 包

常用探测器。

#### MIME

```go
// DetectContentType 监测数据的 MIME 类型
func DetectContentType(data []byte) string
```





## 二、构建

```bash
# iOS
gomobile bind -v -target ios -o "$BASE_DIR/../build/common_utils.xcframework" "$BASE_DIR/../pkg"/*

# Android
gomobile bind -v -target android -androidapi 16 -o "$BASE_DIR/../build/common_utils.aar" "$BASE_DIR/../pkg"/*
```





