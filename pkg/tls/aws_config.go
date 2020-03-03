package aws

type AWSConfig struct {
	Key                     string `json:"key" env:"AWS_CA_KEY"`
	Secret                  string `json:"secret" env:"AWS_CA_SECRET"`
	CertificateAuthorityARN string `json:"ca_arn" env:"AWS_CA_ARN"`
}
