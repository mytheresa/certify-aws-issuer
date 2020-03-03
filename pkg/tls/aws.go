package aws

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"time"

	awsSDK "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/johanbrandhorst/certify"
	"github.com/sirupsen/logrus"

	aws "github.com/robopuff/certify-aws-issuer/pkg/issuer"
	"github.com/robopuff/certify-aws-issuer/pkg/logger"
)

// NewAWSIssuer creates a new instance of AWS issuer
func NewAWSIssuer(cli *acmpca.Client, arn string) (*aws.Issuer, error) {
	return &aws.Issuer{
		Client:                  cli,
		CertificateAuthorityARN: arn,
		TimeToLive:              365,
	}, nil
}

// NewAWSIssuerFromConfig create AWS issuer based on config
func NewAWSIssuerFromConfig(cfg AWSConfig) (*aws.Issuer, error) {
	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}

	awsCfg.Region = endpoints.EuCentral1RegionID
	awsCfg.Credentials = awsSDK.NewStaticCredentialsProvider(cfg.Key, cfg.Secret, "")
	cli := acmpca.New(awsCfg)

	return NewAWSIssuer(cli, cfg.CertificateAuthorityARN)
}

// NewAWSTLSConfig provides a tls config with AWS Private CAs
func NewAWSTLSConfig(awsIssuer *aws.Issuer, options ...AWSOption) (*tls.Config, error) {
	// Default token options
	opts := awsOptions{
		ctx:         context.Background(),
		cache:       certify.NewMemCache(),
		logger:      logger.NewCertifyLogrusMapper(logrus.WithTime(time.Now())),
		renewBefore: (366 * 24) * time.Hour,
	}

	for _, o := range options {
		o.apply(&opts)
	}

	c := &certify.Certify{
		CommonName:  opts.commonName,
		Cache:       opts.cache,
		RenewBefore: opts.renewBefore,
		Logger:      opts.logger,
		Issuer:      awsIssuer,
		CertConfig: &certify.CertConfig{
			SubjectAlternativeNames:   opts.subjectAlternativeNames,
			IPSubjectAlternativeNames: opts.ipSubjectAlternativeNames,
		},
	}

	awsCA, err := awsIssuer.GetCA(opts.ctx)
	if err != nil {
		return nil, err
	}

	cp := x509.NewCertPool()
	cp.AddCert(awsCA)

	return &tls.Config{
		GetCertificate:       c.GetCertificate,
		GetClientCertificate: c.GetClientCertificate,
		RootCAs:              cp,
		ServerName:           opts.commonName,
	}, nil
}
