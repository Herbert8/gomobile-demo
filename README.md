# 使用 gomobile 编写证书链校验 Lib



### 方法说明

```go
// CheckSignatureFrom 判断当前证书是否继承自另一个证书
func (receiver *CertificateWrapper) CheckSignatureFrom(parentCertWrapperObj *CertificateWrapper) error;

// CommonName 获取证书的 CommonName
func (receiver *CertificateWrapper) CommonName() string;

// PublicKeyPEMString 获取证书的 PublicKey，格式为 PEM
func (receiver *CertificateWrapper) PublicKeyPEMString() (string, error);



// NewCertificateWrapperFromPEMData 通过 PEM 格式的数据创建证书
func NewCertificateWrapperFromPEMData(pemData []byte) (*CertificateWrapper, error);

// NewCertificateWrapperFromPEMString 通过 PEM 格式的字符串创建证书
func NewCertificateWrapperFromPEMString(pemStr string) (*CertificateWrapper, error);

// CheckPEMCertSignatureFromParentPEMCertString 通过 PEM 格式字符串判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertString(pemCertStr string, parentPemCertStr string) error;

// CheckPEMCertSignatureFromParentPEMCertData 通过 PEM 格式数据判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertData(pemCertData []byte, parentPemCertData []byte) error;
```





### 构建

```bash
# iOS
gomobile bind -v -target ios -o "./build/cert.xcframework" "./pkg"

# Android
gomobile bind -v -target android -androidapi 32 -o "./build/cert.aar" "./pkg"
```





