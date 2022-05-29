//
//  ViewController.m
//  TestGoLib
//
//  Created by Ba Hongbin on 2022/5/28.
//

#import "ViewController.h"

#import <Cert/Cert.h>

@interface ViewController ()

@end

@implementation ViewController {
    NSString *cert1File;
    NSString *cert2File;
    
}

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    
//    NSLog(@"did load start");
    
    [self prepareFileName];
    
//    NSLog(@"did load finished");
    
}

- (IBAction)btnCheckByObj:(UIButton *)sender {
    [self checkCertWithObj];
}


- (IBAction)btnCheckByData:(UIButton *)sender {
    [self checkCertSimpleWithData];
}


- (IBAction)btnCheckByString:(UIButton *)sender {
    [self checkCertSimpleWithString];
}

- (void)prepareFileName {
    cert1File = [[NSBundle mainBundle] pathForResource:@"1"
                                                          ofType:@"crt"];
    cert2File = [[NSBundle mainBundle] pathForResource:@"2"
                                                          ofType:@"crt"];
}

- (void)checkCertWithObj {
    
    NSData *cert1Data = [NSData dataWithContentsOfFile:cert1File];
    
    NSLog(@"cert1 data size = %ld", cert1Data.length);
    
    
    NSError *err = nil;
//    NSData *cert2Data = [NSData dataWithContentsOfFile:cert2File];
    NSString *cert2Str = [NSString stringWithContentsOfFile:cert2File
                                                   encoding:NSUTF8StringEncoding
                                                      error:&err];
    
    NSLog(@"cert2 str size = %ld", cert2Str.length);
//    NSLog(@"cert2 == %@", cert2Str);
    
    NSError *errCert1 = nil;
//    CertCertificateWrapper *cert1 = CertNewCertificateWrapperFromDERData(cert1Data, &errCert1);
    CertCertificateWrapper *cert1 = CertNewCertificateWrapperFromPEMData(cert1Data, &errCert1);
    
    NSLog(@"cert1 = %@, CommonName = %@", cert1, [cert1 commonName]);
    
    
    NSError *errCert2 = nil;
    CertCertificateWrapper *cert2 = CertNewCertificateWrapperFromPEMString(cert2Str, &errCert2);

    
    NSLog(@"cert2 = %@, CommonName = %@", cert2, [cert2 commonName]);
    
    if (errCert1 == nil && errCert2 == nil) {
        err = nil;
        [cert1 checkSignatureFrom:cert2 error:&err];
        if (err == nil) {
            NSLog(@"OK");
        } else {
            NSLog(@"err = %@", err);
        }

    } else {
        NSLog(@"cert1 err = %@", errCert1);
        NSLog(@"cert2 err = %@", errCert2);
    }
    
}


- (void)checkCertSimpleWithString {
    
    NSError *errCert1 = nil;
    NSError *errCert2 = nil;
    
    NSString *cert1Str = [NSString stringWithContentsOfFile:cert1File
                                                   encoding:NSUTF8StringEncoding
                                                      error:&errCert1];
    NSString *cert2Str = [NSString stringWithContentsOfFile:cert2File
                                                   encoding:NSUTF8StringEncoding
                                                      error:&errCert2];
    
    if (errCert1 == nil && errCert2 == nil) {
        NSError *err = nil;
        CertCheckPEMCertSignatureFromParentPEMCertString(cert1Str, cert2Str, &err);
        if (err == nil) {
            NSLog(@"checkCertSimpleWithString OK");
        } else {
            NSLog(@"checkCertSimpleWithString err = %@", err);
        }

    } else {
        NSLog(@"checkCertSimpleWithString cert1 err = %@", errCert1);
        NSLog(@"checkCertSimpleWithString cert2 err = %@", errCert2);
    }

}


- (void)checkCertSimpleWithData {
    
    NSData *cert1Data = [NSData dataWithContentsOfFile:cert1File];
    NSData *cert2Data = [NSData dataWithContentsOfFile:cert2File];

    NSError *err = nil;
    CertCheckPEMCertSignatureFromParentPEMCertData(cert1Data, cert2Data, &err);
    
    if (err == nil) {
        NSLog(@"checkCertSimpleWithData OK");
    } else {
        NSLog(@"checkCertSimpleWithData err = %@", err);
    }
    
}


@end
