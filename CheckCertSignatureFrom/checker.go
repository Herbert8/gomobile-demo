/*
	证书签发关系验证
	参考《如何验证证书链的签出关系》
	链接：https://studygolang.com/articles/29490
*/
package cert

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// CertificateWrapper 提供一个证书类的包装，用于证书处理。
type CertificateWrapper struct {
	certificate *x509.Certificate
}

// CheckSignatureFrom 判断当前证书是否继承自另一个证书
func (receiver *CertificateWrapper) CheckSignatureFrom(parentCertWrapperObj *CertificateWrapper) error {
	return receiver.certificate.CheckSignatureFrom(parentCertWrapperObj.certificate)
}

// CommonName 获取证书的 CommonName
func (receiver *CertificateWrapper) CommonName() string {
	return receiver.certificate.Subject.CommonName
}

func newCertificateWrapperFromDERData(derData []byte) (*CertificateWrapper, error) {
	cert, err := x509.ParseCertificate(derData)
	if err != nil {
		return nil, err
	}

	return &CertificateWrapper{
		certificate: cert,
	}, nil
}

// NewCertificateWrapperFromPEMData 通过 PEM 格式的数据创建证书
func NewCertificateWrapperFromPEMData(pemData []byte) (*CertificateWrapper, error) {

	block, _ := pem.Decode(pemData)
	if block == nil {
		err := errors.New("ERROR: block of decoded certificate is nil")
		return nil, err
	}

	return newCertificateWrapperFromDERData(block.Bytes)
}

// NewCertificateWrapperFromPEMString 通过 PEM 格式的字符串创建证书
func NewCertificateWrapperFromPEMString(pemStr string) (*CertificateWrapper, error) {
	return NewCertificateWrapperFromPEMData([]byte(pemStr))
}

// CheckPEMCertSignatureFromParentPEMCertString 通过 PEM 格式字符串判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertString(pemCertStr string, parentPemCertStr string) error {
	return CheckPEMCertSignatureFromParentPEMCertData([]byte(pemCertStr), []byte(parentPemCertStr))
}

// CheckPEMCertSignatureFromParentPEMCertData 通过 PEM 格式数据判断证书继承关系
func CheckPEMCertSignatureFromParentPEMCertData(pemCertData []byte, parentPemCertData []byte) error {
	cert, err := NewCertificateWrapperFromPEMData(pemCertData)
	if err != nil {
		return err
	}
	parentCert, err := NewCertificateWrapperFromPEMData(parentPemCertData)
	if err != nil {
		return err
	}
	return cert.CheckSignatureFrom(parentCert)
}
