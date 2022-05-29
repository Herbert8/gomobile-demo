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

type CertificateWrapper struct {
	certificate *x509.Certificate
}

func (receiver *CertificateWrapper) CheckSignatureFrom(parentCertWrapperObj *CertificateWrapper) error {
	return receiver.certificate.CheckSignatureFrom(parentCertWrapperObj.certificate)
}

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

func NewCertificateWrapperFromPEMData(pemData []byte) (*CertificateWrapper, error) {

	block, _ := pem.Decode(pemData)
	if block == nil {
		err := errors.New("ERROR: block of decoded certificate is nil")
		return nil, err
	}

	return newCertificateWrapperFromDERData(block.Bytes)
}

func NewCertificateWrapperFromPEMString(pemStr string) (*CertificateWrapper, error) {
	return NewCertificateWrapperFromPEMData([]byte(pemStr))
}

func CheckPEMCertSignatureFromParentPEMCertString(pemCertStr string, parentPemCertStr string) error {
	return CheckPEMCertSignatureFromParentPEMCertData([]byte(pemCertStr), []byte(parentPemCertStr))
}

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
