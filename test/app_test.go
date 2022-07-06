package test_test

import (
	cert "check_cert/pkg"
	"io/ioutil"
	"log"
	"testing"
)

//func TestCertChecker1(t *testing.T) {
//	pemStr, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/1.crt")
//	derData, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/2.der")
//	cert1, _ := cert.NewCertificateWrapperFromPEMString(string(pemStr))
//	cert2, _ := cert.NewCertificateWrapperFromDERData(derData)
//	if err := cert1.CheckSignatureFrom(cert2); err == nil {
//		t.Log("OK")
//	} else {
//		t.Log(err.Error())
//	}
//}
//
//func TestCertChecker2(t *testing.T) {
//	derData, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/1.der")
//	pemStr, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/2.crt")
//	cert2, _ := cert.NewCertificateWrapperFromPEMString(string(pemStr))
//	cert1, _ := cert.NewCertificateWrapperFromDERData(derData)
//	if err := cert1.CheckSignatureFrom(cert2); err == nil {
//		t.Log("OK")
//	} else {
//		t.Log(err.Error())
//	}
//}

const CERT_PATH = "/Volumes/Data/MyCode/Learning/Golang/gomobile-demo/test/cert_file/"
const CRT1 = CERT_PATH + "1.crt"
const CRT2 = CERT_PATH + "2.crt"

func loadPEMCertStrFromPEMCertFile(pemFileName string) string {
	pemData, _ := ioutil.ReadFile(pemFileName)
	return string(pemData)
}

func loadPEMPubKeyStrFromPEMCertStr(pemCertFile string) string {
	pemCertStr := loadPEMCertStrFromPEMCertFile(pemCertFile)
	certWrapper, _ := cert.NewCertificateWrapperFromPEMString(pemCertStr)
	pubKeyPEMStr, _ := certWrapper.PublicKeyPEMString()
	return pubKeyPEMStr
}

func TestPublicKey(t *testing.T) {
	pbk1Str := loadPEMPubKeyStrFromPEMCertStr(CRT1)
	pbk2Str := loadPEMPubKeyStrFromPEMCertStr(CRT2)
	log.Println(pbk1Str)
	log.Println(pbk2Str)
}
