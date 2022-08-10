package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	errKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	//errNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	errNotRSAPublicKey = errors.New("key is not a valid RSA public key")
)

func loadCertificateFromPEM(pemData []byte) (*x509.Certificate, error) {

	block, _ := pem.Decode(pemData)
	if block == nil {
		err := errors.New("ERROR: block of decoded certificate is nil")
		return nil, err
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, err
}

func extractPublicKeyFromCertificate(cert *x509.Certificate) (*rsa.PublicKey, error) {

	// 提取公钥
	publicKeyDer, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil, err
	}

	pubKeyAny, err := x509.ParsePKIXPublicKey(publicKeyDer)
	if err != nil {
		return nil, err
	}

	var pubKey *rsa.PublicKey
	var ok bool
	if pubKey, ok = pubKeyAny.(*rsa.PublicKey); !ok {
		return nil, errNotRSAPublicKey
	}

	return pubKey, nil
}

// 本方法出自 jwt 库
func parseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, errNotRSAPublicKey
	}

	return pkey, nil
}

// ExtractPublicKeyPEMFromCertificatePEM 获取证书的 PublicKey，格式为 PEM
func ExtractPublicKeyPEMFromCertificatePEM(certPEMData []byte) ([]byte, error) {
	cert, err := loadCertificateFromPEM(certPEMData)
	if err != nil {
		return nil, err
	}

	publicKeyDer, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil, err
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	return pem.EncodeToMemory(&publicKeyBlock), nil
}

// EncryptWithCertificate 使用证书对数据进行加密
func EncryptWithCertificate(certPEMData []byte, msg []byte) ([]byte, error) {
	rsaCert, err := loadCertificateFromPEM(certPEMData)
	if err != nil {
		return nil, err
	}

	pubKey, err := extractPublicKeyFromCertificate(rsaCert)
	if err != nil {
		return nil, err
	}

	encryptedMsg, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
	if err != nil {
		return nil, err
	}
	return encryptedMsg, nil
}

// EncryptWithPublicKey 使用公钥对数据进行加密
func EncryptWithPublicKey(pubKeyPEMData []byte, msg []byte) ([]byte, error) {

	pubKey, err := parseRSAPublicKeyFromPEM(pubKeyPEMData)
	if err != nil {
		return nil, err
	}

	encryptedMsg, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
	if err != nil {
		return nil, err
	}
	return encryptedMsg, nil
}
