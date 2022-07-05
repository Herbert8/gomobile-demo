package test_test

import (
	cert "check_cert"
	"io/ioutil"
	"testing"
)

func TestCertChecker1(t *testing.T) {
	pemStr, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/1.crt")
	derData, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/2.der")
	cert1, _ := cert.NewCertificateWrapperFromPEMString(string(pemStr))
	cert2, _ := cert.NewCertificateWrapperFromDERData(derData)
	if err := cert1.CheckSignatureFrom(cert2); err == nil {
		t.Log("OK")
	} else {
		t.Log(err.Error())
	}
}

func TestCertChecker2(t *testing.T) {
	derData, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/1.der")
	pemStr, _ := ioutil.ReadFile("/Volumes/Data/MyCode/Learning/Golang/CheckCertSignatureFrom/cert_file/2.crt")
	cert2, _ := cert.NewCertificateWrapperFromPEMString(string(pemStr))
	cert1, _ := cert.NewCertificateWrapperFromDERData(derData)
	if err := cert1.CheckSignatureFrom(cert2); err == nil {
		t.Log("OK")
	} else {
		t.Log(err.Error())
	}
}
