package aws

import (
	"context"
	"net"
	"time"

	"github.com/johanbrandhorst/certify"
)

type awsOptions struct {
	ctx                       context.Context
	commonName                string
	subjectAlternativeNames   []string
	ipSubjectAlternativeNames []net.IP
	renewBefore               time.Duration
	logger                    certify.Logger
	cache                     certify.Cache
}

// AWSOption Options for the token
type AWSOption interface {
	apply(*awsOptions)
}

type funcAWSOption struct {
	f func(*awsOptions)
}

func (fdo funcAWSOption) apply(do *awsOptions) {
	fdo.f(do)
}

func newFuncAWSOption(f func(*awsOptions)) *funcAWSOption {
	return &funcAWSOption{
		f: f,
	}
}

// Context add context to options
func Context(ctx context.Context) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.ctx = ctx
	})
}

// CommonName add subject common name to options
func CommonName(cn string) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.commonName = cn
	})
}

// SubjectAlternativeNames add alternative subject names to options
func SubjectAlternativeNames(n []string) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.subjectAlternativeNames = n
	})
}

// IpSubjectAlternativeNames add alternative subject ip names to options
func IpSubjectAlternativeNames(i []net.IP) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.ipSubjectAlternativeNames = i
	})
}

// RenewBefore add time duration to option
func RenewBefore(t time.Duration) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.renewBefore = t
	})
}

// Logger add logger to options
func Logger(l certify.Logger) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.logger = l
	})
}

// Cache add cache to options
func Cache(c certify.Cache) AWSOption {
	return newFuncAWSOption(func(o *awsOptions) {
		o.cache = c
	})
}
